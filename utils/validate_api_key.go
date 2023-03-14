package utils

import "regexp"

func ValidateApiKey(apiKey string) bool {
	pattern := regexp.MustCompile(`^sk-[a-zA-Z0-9]{48}$`)
	return pattern.MatchString(apiKey)
}
