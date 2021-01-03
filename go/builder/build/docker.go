package build

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

/*
interact with docker lib to perform docker actions
*/

func Docker(ctx context.Context, param BuildParameters) {
	fmt.Println("Download and save docker images ...")
	pathToImageList := "./dist/imageList"
	ExtractImageFromList(pathToImageList)
}

func ExtractImageFromList(pathToImageList string) {
	if _, err := os.Stat(pathToImageList); os.IsExist(err) {
		log.Fatal("imagesList file does not exist")
	}

	file, err := os.Open(pathToImageList)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	path := "./dist/images"
	err = os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal(err)
	}
	var waitGroup sync.WaitGroup
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		image := scanner.Text()
		waitGroup.Add(1)
		go pullAndSave(image, &waitGroup)
		fmt.Println("Save image:", image, "...")
	}
	waitGroup.Wait()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// pull pulls an image using "docker pull" command that lets us take advantage of its cached
// credentials for multiple docker registries
func pullAndSave(image string, waitGroup *sync.WaitGroup) {
	cmd := exec.Command("docker", "pull", image)
	out, err := cmd.CombinedOutput()
	if err != nil {
		waitGroup.Done()
		log.Fatal("Image:", image, string(out))
		log.Fatal(err)
	}
	r := regexp.MustCompile(`Digest: (.*)\n`)
	digest := r.FindStringSubmatch(string(out))
	digestValue := strings.Split(digest[1], ":")[1]

	if _, err := os.Stat("dist/images/" + digestValue + ".tar"); os.IsNotExist(err) {
		cmd = exec.Command("docker", "save", image)
		var buf bytes.Buffer
		cmd.Stdout = &buf
		err = cmd.Run()
		if err != nil {
			waitGroup.Done()
			log.Println(cmd.Stderr)
			log.Fatal(err)
		}
		//, ">", cwd+"/dist/images/"+digestValue+".tar"
		createTarballFile("dist/images", digestValue+".tar", buf.Bytes())
	}
	fmt.Println("Save image:", image, "=>", digest[1])
	waitGroup.Done()
}

func createTarballFile(path string, filename string, content []byte) error {
	tarFilePath := filepath.Join(path, filename)
	readmeFile, err := os.Create(tarFilePath)
	defer readmeFile.Close()
	if err != nil {
		errMsg := fmt.Errorf("%s", err)
		fmt.Println(errMsg)
		return err
	}

	err = ioutil.WriteFile(tarFilePath, content, 0644)
	return err
}
