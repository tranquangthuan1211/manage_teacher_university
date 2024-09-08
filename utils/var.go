package utils

import (
	"os"
)

func getenv(key, fallBack string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallBack
	}

	return value
}

var SECRET_KEY = []byte(getenv("SECRET_KEY", "tranquanthuan132@gmail.com"))

var PORT = getenv("PORT", "8080")
