package gen

import (
	"fmt"
	"helm.sh/helm/v3/pkg/chartutil"
)

func Helm(param GenParameters) {
	fmt.Println("create sample helm chart", param)
	path, err := CreateFolder(param, "charts")
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

	_, err = chartutil.Create(param.Name, path)
	if err != nil {
		errMsg := fmt.Errorf("%s", err)
		fmt.Println(errMsg)
	}
}
