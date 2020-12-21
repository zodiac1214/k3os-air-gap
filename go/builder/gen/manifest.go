package gen

import (
	"fmt"
	"path/filepath"
)

func Manifest(param GenParameters) {
	fmt.Println("create manifest file", param)
	const content = `
apiVersion: zodiac1214/k3os
kind: Manifest
metadata:
  name: ThisIsTheName
`
	CreateTextFile(filepath.Join(param.Path, param.Name), "manifest.yaml", content)
}
