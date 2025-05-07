package services

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/farpat/go-url-shortener/internal/utils/url"
)

func GenerateSlug(inputURL string) string {
	hasher := md5.New()
	normalizedUrl, _ := url.NormalizeURL(inputURL)
	normalizedUrlAsBytes := []byte(normalizedUrl)
	hasher.Write(normalizedUrlAsBytes)
	return hex.EncodeToString(hasher.Sum(nil))
}
