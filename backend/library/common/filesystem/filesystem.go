package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// check if file/directory from this path exist
func PathExist(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) || err != nil {
		return false
	}

	return true
}

// get content from file
func FileGetContent(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	dataString := string(data)
	return dataString, nil
}

// put content into file
func FilePutContent(filePath string, data string) error {
	// convert slashes first
	filePath = filepath.ToSlash(filePath)
	pathSlices := strings.Split(filePath, "/")
	total := len(pathSlices)

	// get directory path
	dirPathSlices := pathSlices[0 : total-1]
	dirPath := strings.Join(dirPathSlices, "/")

	// check if there is slash prefix especially in UNIX system
	if filePath[0:1] == "/" {
		dirPath = "/" + dirPath
	}

	// check if directory exist and create if not exist
	if !PathExist(dirPath) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory in path: %s. Detail: %v", dirPath, err)
		}
	}

	// create file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file in path: %s. Detail: %v", filePath, err)
	}
	defer file.Close()

	// write to file
	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("failed to write data to file in path: %s. Detail: %v", filePath, err)
	}

	// return
	return nil
}
