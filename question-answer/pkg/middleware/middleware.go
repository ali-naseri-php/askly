package middleware

import (
	"errors"
	"strings"
)

// ExtractUserIDFromToken استخراج userID از توکن
func ExtractUserIDFromToken(token string) (string, error) {
	if token == "" {
		return "", errors.New("invalid token")
	}
	parts := strings.Split(token, ":")
	if len(parts) != 2 {
		return "", errors.New("invalid token format")
	}
	return parts[1], nil
}
