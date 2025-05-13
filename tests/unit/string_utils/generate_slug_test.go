package string_utils_test

import (
	"testing"

	"github.com/farpat/go-url-shortener/internal/services/string_utils"
	"github.com/stretchr/testify/assert"
)

func TestSlugsAreEquals(t *testing.T) {
	// ARRANGE
	url1 := "https://www.google.com"
	url2 := "https://www.google.com"

	// ACT
	result1 := string_utils.GenerateSlug(url1)
	result2 := string_utils.GenerateSlug(url2)

	// ASSERT
	assert.Equal(t, result1, result2)
}

func TestSlugsAreNotEquals(t *testing.T) {
	// ARRANGE
	url1 := "https://www.google.com"
	url2 := "https://www.github.com"

	// ACT
	result1 := string_utils.GenerateSlug(url1)
	result2 := string_utils.GenerateSlug(url2)

	// ASSERT
	assert.NotEqual(t, result1, result2)
}

func TestSlugsIsProperlyGenerated(t *testing.T) {
	// ARRANGE
	url := "https://www.google.com"

	// ACT
	result := string_utils.GenerateSlug(url)

	// ASSERT
	expected := "263c66c0dcbc045e38f2a5fa5f47341f"
	assert.Equal(t, expected, result)
}
