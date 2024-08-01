package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKeyToken(headers http.Header) (string, error) {
	apiKeyHeader := headers.Get("Authorization")
	if apiKeyHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(apiKeyHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}
