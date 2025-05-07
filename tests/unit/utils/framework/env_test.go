package framework_test

import (
	"os"
	"strings"
	"testing"

	"github.com/farpat/go-url-shortener/internal/utils/framework"
	"github.com/stretchr/testify/assert"
)

func TestDefaultValue(t *testing.T) {
	// ARRANGE
	key := "TEST_UNEXISTING_ENV_VARIABLE"
	defaultValue := "default"

	// ACT
	result := framework.Env(key, defaultValue)

	// ASSERT
	assert.Equal(t, defaultValue, result)
}

func TestValue(t *testing.T) {
	key, value := extractFirstEnvVariable()

	// ACT
	result := framework.Env(key, "not used")

	// ASSERT
	assert.Equal(t, value, result)
	assert.NotEqual(t, "not used", result)
}

func extractFirstEnvVariable() (key string, value string) {
	exportedVariable := os.Environ()[0]

	equalsPosition := strings.Index(exportedVariable, "=")
	key = exportedVariable[:equalsPosition]
	value = exportedVariable[equalsPosition+1:]

	return key, value
}
