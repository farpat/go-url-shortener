package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func ProjectPath(relativePath string) string {
	currentPath, _ := os.Getwd()
	if strings.HasSuffix(currentPath, "/public") {
		currentPath = filepath.Dir(currentPath)
	}

	return currentPath + "/" + relativePath
}
