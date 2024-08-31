package utils

import (
	"errors"
	"strings"
)

// GenerateUsername : Function for generate username
// email : Email of account
func GenerateUsername(email string) (string, error) {
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		return "", errors.New("failed to generate username")
	}

	username := "acc-" + email[:atIndex]
	return username, nil
}
