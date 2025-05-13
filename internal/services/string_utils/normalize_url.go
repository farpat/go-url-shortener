package string_utils

import (
	"regexp"
	"strings"

	internalErrors "github.com/farpat/go-url-shortener/internal/errors"
)

func NormalizeURL(urlToNormalize string) (string, error) {
	stringToHandle := strings.ToLower(urlToNormalize)
	if regex, err := regexp.Compile(`^(https?:\/\/)?(www\.)?([\w\-]+\.\w{1,3}\/?(.*))$`); err == nil {
		matches := regex.FindStringSubmatch(stringToHandle)
		if len(matches) == 0 {
			return "", &internalErrors.InvalidUrlError{Url: urlToNormalize}
		}

		stringToHandle = matches[3]
	}

	normalizedString := ""

	for _, char := range stringToHandle {
		if char >= 'a' && char <= 'z' || char >= '0' && char <= '9' {
			normalizedString += string(char)
		} else {
			normalizedString += "-"
		}
	}

	return strings.Trim(normalizedString, "-"), nil
}
