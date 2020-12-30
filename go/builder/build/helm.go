package build

import (
	"context"
	"embed"
	"fmt"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/engine"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

/*
interact with helm lib
*/

func Helm(ctx context.Context, param BuildParameters) {
	err := extractSystemCharts()
	if err != nil {
		log.Fatal(err)
	}

	err = forceImagePullPolicyToLocal()
	if err != nil {
		log.Fatal(err)
	}

	extractImagesFromRenderedCharts(param)
}

func forceImagePullPolicyToLocal() error {
	err := filepath.Walk("dist/charts", findAndReplace)
	return err
}

func findAndReplace(path string, fi fs.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil //
	}

	matched, err := filepath.Match("*.yaml", fi.Name())

	if err != nil {
		panic(err)
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		//fmt.Println(string(read))
		fmt.Println(path)

		newContents := strings.Replace(string(read), "imagePullPolicy: Always", "imagePullPolicy: Never", -1)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}

	}

	return nil
}

func extractImagesFromRenderedCharts(param BuildParameters) {
	fmt.Println("Render helm charts ...")
	pathToCharts := param.Path + "/charts"
	files, err := ioutil.ReadDir(pathToCharts)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		name := f.Name()
		path := pathToCharts + "/" + name
		if isValidChart, _ := chartutil.IsChartDir(path); !isValidChart {
			errPrint := fmt.Errorf("%s", "Failed to process chart at: "+path)
			fmt.Println(errPrint)
			os.Exit(1)
		}

		fmt.Println("process chart:", path)
		chart, err := loader.Load(path)
		if err != nil {
			errPrint := fmt.Errorf("%s", "Failed to load chart : "+path)
			fmt.Println(errPrint)
			os.Exit(1)
		}
		values, err := chartutil.ReadValuesFile(path + "/values.yaml")
		renderValues, err := chartutil.ToRenderValues(chart, values, chartutil.ReleaseOptions{}, chartutil.DefaultCapabilities)
		renderedTemplates, err := engine.Render(chart, renderValues)
		if err != nil {
			errPrint := fmt.Errorf("%s", "Failed to render chart : "+path)
			fmt.Println(errPrint)
			log.Fatal(err)
		}

		result := make(map[string]string)
		for k, v := range renderedTemplates {
			filename := filepath.Base(k)
			// Render only Kubernetes resources skipping internal Helm
			// files and files that begin with underscore which are not
			// expected to output a Kubernetes spec.
			if filename == "NOTES.txt" || strings.HasPrefix(filename, "_") {
				continue
			}
			result[k] = v
			r := regexp.MustCompile(`image: (.*)\n`)
			foundStrings := r.FindStringSubmatch(v)
			if len(foundStrings) == 2 {
				content := strings.Replace(foundStrings[1], `"`, "", -1) + "\n"
				if len(content) > 1 {
					fmt.Println("Extract image from rendered chart, image found:", content)
					appendToFile("./dist", "imageList", content)
				}
			}
		}
	}
}

//go:embed system-charts/*
var SysChartsFiles embed.FS

func extractSystemCharts() error {
	fmt.Println("Extract system chart ...")
	err := ExtractBundledDirectory("system-charts", SysChartsFiles, "charts")
	if err != nil {
		return err
	}
	os.Rename("./dist/system-charts", "./dist/charts2")
	return nil
}

func appendToFile(path string, filename string, content string) error {
	f, err := os.OpenFile(filepath.Join(path, filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		errMsg := fmt.Errorf("%s", err)
		fmt.Println(errMsg)
		return err
	}
	return nil
}
