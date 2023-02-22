package util

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ListAllFolder(path string) ([]string, error) {
	info, error := ioutil.ReadDir(path)
	if error != nil {
		return nil, error
	}

	var folders []string
	for _, f := range info {
		if f.IsDir() {
			folders = append(folders, f.Name())
		}
	}

	return folders, nil
}

func GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func ListAllFiles(path string) ([]string, error) {
	info, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, f := range info {
		if !f.IsDir() {
			files = append(files, f.Name())
		}
	}

	return files, nil
}

func GetFilenameFromPath(filePath string) string {
	if filePath == "" {
		return ""
	}
	return filepath.Base(filePath)
}

func GetFilenameWithoutExt(filename string) string {
	filename = GetFilenameFromPath(filename)
	if strings.HasPrefix(filename, ".") {
		return filename
	}
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}
