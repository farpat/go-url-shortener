package url

import (
	"errors"
	"regexp"
	"strings"
)

// NormalizeURL normalizes a URL to be used in a slug
func NormalizeURL(urlToNormalize string) (string, error) {
	stringToHandle := strings.ToLower(urlToNormalize)
	if regex, err := regexp.Compile(`^(https?:\/\/)?(www\.)?([\w\-]+\.\w{1,3}\/?(.*))$`); err == nil {
		matches := regex.FindStringSubmatch(stringToHandle)
		if len(matches) == 0 {
			return "", errors.New("URL invalide")
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
