package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileUtil_GetCurrentFolder(t *testing.T) {
	assert := assert.New(t)

	currentFolderPath := NewFileUtil().GetCurrentFolder()
	assert.NotEmpty(currentFolderPath)
	// Current folder name should be "util"
	assert.Regexp(`util$`, currentFolderPath)
}

func TestFileUtil_GetAbsolutePathOf(t *testing.T) {
	f := NewFileUtil()

	// Arrange
	relativePath := "../sample"

	// Act
	absolutePath := f.GetAbsolutePathOf(relativePath)

	// Assert
	assert := assert.New(t)
	assert.NotEmpty(absolutePath)
	assert.Regexp(`sample$`, absolutePath)

	relativePath2 := "./file_util_test.go"
	absolutePath2 := f.GetAbsolutePathOf(relativePath2)
	assert.Regexp(`file_util_test.go$`, absolutePath2)

	nilRelativePathAbsolute := f.GetAbsolutePathOf("")
	assert.Equal(nilRelativePathAbsolute, f.GetCurrentFolder())
}
