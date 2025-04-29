package utils

import "strings"

func NormalizeString(stringToNormalize string) string {
	stringToHandle := strings.ToLower(strings.TrimSpace(stringToNormalize))
	stringToHandle = strings.TrimPrefix(stringToHandle, "https://")
	stringToHandle = strings.TrimPrefix(stringToHandle, "http://")
	stringToHandle = strings.TrimPrefix(stringToHandle, "www.")

	normalizedString := ""

	for _, char := range stringToHandle {
		if char >= 'a' && char <= 'z' || char >= '0' && char <= '9' {
			normalizedString += string(char)
		} else {
			normalizedString += "-"
		}
	}

	return strings.Trim(normalizedString, "-")
}
