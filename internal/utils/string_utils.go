package utils

import "strings"

func NormalizeString(stringToNormalize string) string {
	stringToHandle := strings.ToLower(stringToNormalize)
	stringToHandle = strings.TrimSpace(stringToHandle)
	if strings.HasPrefix(stringToHandle, "https://") {
		stringToHandle = strings.Replace(stringToHandle, "https://", "", 1)
	}
	if strings.HasPrefix(stringToHandle, "http://") {
		stringToHandle = strings.Replace(stringToHandle, "http://", "", 1)
	}
	if strings.HasPrefix(stringToHandle, "www.") {
		stringToHandle = strings.Replace(stringToHandle, "www.", "", 1)
	}

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
