package utils

import (
	"testing"

	"github.com/farpat/go-url-shortener/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeString(t *testing.T) {
	// ARRANGE
	stringToNormalizes := map[string]string{
		"https://www.google.com":       "google-com",
		"https://github.com":           "github-com",
		"https://golang.org/":          "golang-org",
		"https://golang.org/titi/toto": "golang-org-titi-toto",
	}

	// ACT
	for stringToNormalize, expected := range stringToNormalizes {
		result := utils.NormalizeString(stringToNormalize)

		// ASSERT
		assert.Equal(t, expected, result)
	}
}
