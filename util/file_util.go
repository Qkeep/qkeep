package util

import (
	"os"
	"path"
	"path/filepath"
)

type FileUtil struct{}

func NewFileUtil() *FileUtil {
	return &FileUtil{}
}

// GetCurrentFolder returns the absolute path of the current folder.
//
// Example:
//
//	absPath := NewFileUtil().GetCurrentFolder()
//	fmt.Println(absPath) // /home/user
func (f *FileUtil) GetCurrentFolder() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

// GetAbsolutePathOf returns the absolute path of the given relative path
// relative to the current folder.
//
// Example: if the current folder is /home/user and the relative path is ../test,
// the absolute path will be /home/test.
//
// @param relativePath the relative path to the current folder
//
// @return the absolute path of the given relative path
//
// Example:
//
//	absPath := NewFileUtil().GetAbsolutePathOf("../test") // current folder is /home/user
//	fmt.Println(absPath) // /home/test
func (f *FileUtil) GetAbsolutePathOf(relativePath string) string {
	currentFolderAbsPath := f.GetCurrentFolder()
	absolutePath, err := filepath.Abs(path.Join(currentFolderAbsPath, relativePath))
	if err != nil {
		panic(err)
	}

	return absolutePath
}
