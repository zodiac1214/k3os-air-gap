package gen

import (
	"fmt"
	"path/filepath"
)

func Manifest(param GenParameters) {

	const content = `
apiVersion: zodiac1214/kcap
kind: Manifest
metadata:
  name: ThisIsTheName
`
	path := filepath.Join(param.Path, param.Name)
	CreateTextFile(path, "manifest.yaml", content)
	fmt.Println("create manifest file: ./" + path + "/manifest.yaml")
}
