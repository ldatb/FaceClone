package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func CreateAvatarUrl(filename string) (string, error) {
	// Load url from .env
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	// Get image url
	URL := os.Getenv("APP_URL") + "/users/avatar/" + filename

	return URL, nil
}

func Find(array []string, item string) bool {
	for _, i := range array {
		if i == item {
			return true
		}
	}
	return false
}

func FindAndDelete(array []string, item string) []string {
	index := 0
	for _, i := range array {
		if i != item {
			array[index] = i
			index++
		}
	}

	return array[:index]
}