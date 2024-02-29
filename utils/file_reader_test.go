package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFromCorrectFilePath(t *testing.T) {
	filePath := "../person.json"

	fileContents, err := ReadFromJSONFile(filePath)

	assert.Nil(t, err)
	assert.Contains(t, fileContents, "king")
}

func TestReadFileContentFromWrongFilePath(t *testing.T) {
	filePath := "../../person.json"

	fileContents, err := ReadFromJSONFile(filePath)

	assert.NotNil(t, err)
	assert.Empty(t, fileContents)
}

func TestReadEmptyJSONFlie(t *testing.T) {
	filePath := "../empty.person.json"

	fileContents, err := ReadFromJSONFile(filePath)

	assert.NotNil(t, err)
	assert.True(t, fileContents == "")
}
