package framework

import "os"

// Retrieves an environment variable, return a default value if the variable is not set
func Env(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}
	return value
}
