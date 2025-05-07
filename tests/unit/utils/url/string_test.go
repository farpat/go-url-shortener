package url_test

import (
	"testing"

	"github.com/farpat/go-url-shortener/internal/utils/url"
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
		result, err := url.NormalizeString(stringToNormalize)

		// ASSERT
		assert.Equal(t, expected, result)
		assert.NoError(t, err)
	}
}

func TestNormalizeBadString(t *testing.T) {
	// ARRANGE
	invalidURL := "https://"

	// ACT
	result, err := url.NormalizeString(invalidURL)

	// ASSERT
	assert.Equal(t, "", result)
	assert.Error(t, err)
}
