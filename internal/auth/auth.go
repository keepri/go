package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers *http.Header) (string, error) {
	header := headers.Get("Authorization")
	if header == "" {
		return "", errors.New("no auth header found")
	}

	split := strings.Split(header, " ")
	if len(split) != 2 {
		return "", errors.New("malformed auth header")
	}

	if strings.Compare(split[0], "ApiKey") != 0 {
		return "", errors.New("malformed first part of auth header")
	}

	return split[1], nil
}
