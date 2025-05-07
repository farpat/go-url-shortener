package utils

import (
	"testing"

	"github.com/farpat/go-url-shortener/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeGoodString(t *testing.T) {
	// ARRANGE
	stringToNormalizes := map[string]string{
		"https://www.google.com":       "google-com",
		"https://github.com":           "github-com",
		"https://golang.org/":          "golang-org",
		"https://golang.org/titi/toto": "golang-org-titi-toto",
		"www.golang.org/titi/toto":     "golang-org-titi-toto",
		"www.GoLang.org/titi/toto":     "golang-org-titi-toto",
	}

	// ACT
	for stringToNormalize, expected := range stringToNormalizes {
		result, err := utils.NormalizeString(stringToNormalize)

		// ASSERT
		assert.Equal(t, expected, result)
		assert.NoError(t, err)
	}
}

func TestNormalizeBadString(t *testing.T) {
	// ARRANGE
	stringToNormalize := "bad_url"

	// ACT
	result, err := utils.NormalizeString(stringToNormalize)

	// ASSERT
	assert.Equal(t, "", result)
	assert.Error(t, err)
}
