package utils

import (
	"errors"
	"io"
	"os"
)

func ReadFromJSONFile(path string) (string, error) {

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	fileContents, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	if string(fileContents) == "" {
		return "", errors.New("empty json file")
	}

	return string(fileContents), nil
}
