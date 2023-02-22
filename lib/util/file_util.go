package util

import (
	"io/ioutil"
	"log"
	"os"
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
	return filePath[len(filePath)-len(GetFilename(filePath)):]
}

func GetFilename(filePath string) string {
	var filename string
	for i := len(filePath) - 1; i >= 0; i-- {
		if string(filePath[i]) == "/" {
			break
		}
		filename = string(filePath[i]) + filename
	}
	return filename
}

func GetFilenameWithoutExt(filename string) string {
	fileName := GetFilenameFromPath(filename)
	return fileName[:len(fileName)-len(GetExt(fileName))-1]
}

func GetExt(filename string) string {
	lastDotIndex := -1
	for i := len(filename) - 1; i >= 0; i-- {
		if string(filename[i]) == "." {
			lastDotIndex = i
			break
		}
	}
	return filename[lastDotIndex+1:]
}
