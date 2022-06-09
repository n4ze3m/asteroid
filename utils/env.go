package utils

import "os"

func GetEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return ""
}
