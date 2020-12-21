package cli

import (
	"fmt"
	"github.com/zodiac1214/go/builder/gen"
	"os"
)

func RunGen(param gen.GenParameters) {
	fullPath := param.Path + "/" + param.Name
	err := os.MkdirAll(fullPath, 0755)
	if err != nil {
		errPrint := fmt.Errorf("%s", "Failed to create output directory: "+fullPath)
		fmt.Println(errPrint)
		os.Exit(1)
	}
	gen.Helm(param)
	gen.Kubernetes(param)
	gen.Manifest(param)
	const readmeContent = `
# Welcome to k3os air gap project
`
	err = gen.CreateReadmeFile(fullPath, readmeContent)
	if err != nil {
		errPrint := fmt.Errorf("%s", "Failed to create readme.md: "+fullPath)
		fmt.Println(errPrint)
		os.Exit(1)
	}
}
