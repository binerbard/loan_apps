package utils

import (
	"time"

	"math/rand"

	"loan_apps/config/setting"
)

// ExtractToken extracts the token from the Authorization header.
func ExtractToken(token string) string {
	return token[len("Bearer "):]
}

// GenerateRandomString generates a random string of the specified length.
// If length is 0, it defaults to 10.
func GenerateRandomString(length int) (string, error) {
	if length == 0 {
		length = 10
	}
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b), nil
}

// GenerateExpiredTime generates the expiration time for a token.
// It adds the specified number of minutes to the current time.
func GenerateExpiredTime(expiredTime int) time.Time {
	return time.Now().Add(time.Minute * time.Duration(expiredTime))
}

// ExpiredToken generates the expiration time for a token using the
// value of setting.ExpiredToken.
func ExpiredToken() time.Time {
	return GenerateExpiredTime(setting.ExpiredToken)
}

// TokenGenerate generates a random token of the length specified by
// setting.TokenLength.
func TokenGenerate() string {
	token, _ := GenerateRandomString(setting.TokenLength)
	return token
}
