package framework_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/farpat/go-url-shortener/internal/utils/framework"
	"github.com/stretchr/testify/assert"
)

func TestProjectPath(t *testing.T) {
	// ARRANGE
	relativePath := "public"

	// ACT
	result := framework.ProjectPath(relativePath)

	// ASSERT
	currentPath, _ := os.Getwd()
	// Remove tests/unit/utils from currentPath
	expected := filepath.Clean(currentPath+"/../../../../") + "/" + relativePath
	assert.Equal(t, expected, result)
}
