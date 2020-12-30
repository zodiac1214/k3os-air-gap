package build

import (
	"context"
	"fmt"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/engine"
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

			// set imagePullPolicy (TODO)
			r = regexp.MustCompile(`imagePullPolicy: (.*)\n`)
			foundStrings = r.FindStringSubmatch(v)
			//if len(foundStrings) == 2 {
			//	content := strings.Replace(foundStrings[1], `"`, "", -1) + "\n"
			//	if len(content) > 1 {
			//		fmt.Println("Extract image from rendered chart, image found:", content)
			//		appendToFile("./dist", "imageList", content)
			//	}
			//}
		}
	}
}

func extractSystemCharts() error {
	fmt.Println("TODO: extract system chart")
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
