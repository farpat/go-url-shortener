package services

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/farpat/go-url-shortener/internal/utils"
)

func GenerateSlug(url string) string {
	hasher := md5.New()
	normalizedUrl := utils.NormalizeString(url)
	normalizedUrlAsBytes := []byte(normalizedUrl)
	hasher.Write(normalizedUrlAsBytes)
	return hex.EncodeToString(hasher.Sum(nil))
}
