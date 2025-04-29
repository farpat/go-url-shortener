package utils

import (
	"os"
	"strings"
)

const projectName = "go-url-shortener"

func ProjectPath(relativePath string) string {
	currentPath, _ := os.Getwd()
	suffixPosition := strings.LastIndex(currentPath, projectName)
	if suffixPosition == -1 {
		panic("Current directory is not part of the project")
	}
	prefix := strings.Trim(currentPath[:suffixPosition], "/")

	return "/" + prefix + "/" + projectName + "/" + relativePath
}
