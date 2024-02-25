package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func ReadFromJSONFile(path string) (string, error)  {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	defer file.Close()

	fileContents, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	if string(fileContents) == "" {
		return "", errors.New("empty json file")
	}

	return string(fileContents), nil
}	