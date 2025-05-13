package string_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateSlug(inputURL string) string {
	hasher := md5.New()
	normalizedUrl, _ := NormalizeURL(inputURL)
	normalizedUrlAsBytes := []byte(normalizedUrl)
	hasher.Write(normalizedUrlAsBytes)
	return hex.EncodeToString(hasher.Sum(nil))
}
