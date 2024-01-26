package file

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(filename string) ([]byte, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("File reading error: %v", err)
	}
	return fileContent, nil
}
