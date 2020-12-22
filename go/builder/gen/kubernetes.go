package gen

import (
	"fmt"
)

func Kubernetes(param GenParameters) {
	path, err := CreateFolder(param, "kubernetes")
	if err != nil {
		errMsg := fmt.Errorf("%s", err)
		fmt.Println(errMsg)
	}
	fmt.Println("create sample valina k8s yaml files:", path)

	const readmeContent = `
# This is a Readme file HELM
`
	err = CreateTextFile(path, "README.md", readmeContent)
	if err != nil {
		errMsg := fmt.Errorf("%s", err)
		fmt.Println(errMsg)
	}
}
