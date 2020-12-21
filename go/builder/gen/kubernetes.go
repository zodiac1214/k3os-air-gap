package gen

import (
	"fmt"
)

func Kubernetes(param GenParameters) {
	fmt.Println("create sample valina k8s yaml files", param)
	path, err := CreateFolder(param, "kubernetes")
	if err != nil {
		errMsg := fmt.Errorf("%s", err)
		fmt.Println(errMsg)
	}

	const readmeContent = `
# This is a Readme file HELM
`
	err = CreateTextFile(path, "README.md", readmeContent)
	if err != nil {
		errMsg := fmt.Errorf("%s", err)
		fmt.Println(errMsg)
	}
}
