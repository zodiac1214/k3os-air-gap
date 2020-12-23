package build

import (
	"context"
	"embed"
	"fmt"
	"io"
	"log"
	"os"
)

/*
interact with hashicorp's packer lib to create machine images
*/

func Packer(ctx context.Context, param BuildParameters) {
	fmt.Println("Extract packer files to dist folder ...")
	_ = extractBundledDirectory("packer")
	fmt.Println("Building VM image with packer (TODO)")
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
