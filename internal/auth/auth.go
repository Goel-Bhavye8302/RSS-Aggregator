package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Will return an API Key from the headers of the http request
// Example
// Authorization : ApiKey (insert api key here)
func GetApiKey(headers http.Header) (string, error){ 
	key := headers.Get("Authorization")
	if key == "" {
		return  "", errors.New("no API key found")
	}

	vals := strings.Split(key, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return  "", errors.New("malformed Auth Header")
	}

	return vals[1], nil
}