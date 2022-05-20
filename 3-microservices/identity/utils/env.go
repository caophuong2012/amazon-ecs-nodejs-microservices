package utils

import "os"

func GetWithDefault(key string, defaultValue string) string {
	s := os.Getenv(key)

	if s == "" {
		return defaultValue
	}
	return s
}
