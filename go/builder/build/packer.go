package build

import (
	"bufio"
	"context"
	"embed"
	"fmt"
	"log"
)

/*
interact with hashicorp's packer lib to create machine images
*/

func Packer(ctx context.Context, param BuildParameters) {
	fmt.Println("build packer")

	//todo: this just shows how to pack static files in go binary

	file, _ := PackerFiles.Open("packer/system-images.list")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		image := scanner.Text()
		out, err := pullAndSave(image)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(out)
	}

}

//go:embed packer/*
var PackerFiles embed.FS
