package build

import (
	"bufio"
	"context"
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

/*
interact with hashicorp's packer lib to create machine images
*/

func Packer(ctx context.Context, param BuildParameters) {
	fmt.Println("Extract packer files to dist folder ...")
	_ = extractBundledDirectory("packer")

	fmt.Println("Download k3s air gap images ...")
	if _, err := os.Stat("dist/k3s-airgap-images-amd64.tar"); os.IsNotExist(err) {
		fileUrl := "https://github.com/k3s-io/k3s/releases/download/v1.18.9%2Bk3s1/k3s-airgap-images-amd64.tar"
		err := downloadFile("dist/k3s-airgap-images-amd64.tar", fileUrl)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded: " + fileUrl)
	}

	fmt.Println("Building VM image with packer ...")
	var cmd *exec.Cmd
	cmd = exec.Command("packer", "build", "--force", "packer.json")
	cmd.Dir = "./dist/packer/" + param.ImageType
	stdout, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
}

//go:embed packer/*
var PackerFiles embed.FS

func extractBundledDirectory(path string) error {
	directories, err := PackerFiles.ReadDir(path)
	if err != nil {
		log.Fatal("Failed to extract packer", err.Error())
	}
	for _, directory := range directories {
		fullPath := path + "/" + directory.Name()
		fmt.Println("Extracting: ", fullPath)
		if directory.IsDir() {
			err := CreateIfNotExists("dist/"+fullPath, 0755)
			if err != nil {
				log.Fatal("Failed to create folder in dist", err.Error())
			}
			extractBundledDirectory(fullPath)
		} else {
			file, _ := PackerFiles.Open(fullPath)
			out, err := os.Create("dist/" + fullPath)
			if err != nil {
				return err
			}
			_, err = io.Copy(out, file)
		}
	}
	return nil
}

// downloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
