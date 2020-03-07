package utils

import (
	"os"
	"path/filepath"
)

// GetExecutableDir gets the executable directory
func GetExecutableDir() (exPath string, err error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	exPath = filepath.Dir(ex)
	return exPath, nil
}