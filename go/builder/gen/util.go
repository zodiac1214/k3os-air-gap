package gen

import (
	"fmt"
	"os"
	"path/filepath"
)

type GenParameters struct {
	Name  string
	Path  string
	Force bool
}

func CreateFolder(param GenParameters, subFolder string) (string, error) {
	path := param.Path + "/" + param.Name + "/" + subFolder
	err := os.MkdirAll(path, 0755)
	return path, err
}

func CreateTextFile(path string, filename string, content string) error {
	readmeFile, err := os.Create(filepath.Join(path, filename))
	defer readmeFile.Close()
	if err != nil {
		errMsg := fmt.Errorf("%s", err)
		fmt.Println(errMsg)
		return err
	}

	_, err = readmeFile.WriteString(content)
	if err != nil {
		errMsg := fmt.Errorf("%s", err)
		fmt.Println(errMsg)
		return err
	}
	return nil
}
