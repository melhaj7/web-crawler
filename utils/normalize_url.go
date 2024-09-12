package utils

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURl string) (string, error) {
	parsedURL, err := url.Parse(rawURl)
	if err != nil {
		return "", err
	}

	normalized := parsedURL.Host + parsedURL.Path

	if parsedURL.RawQuery != "" {
		normalized += "?" + parsedURL.RawQuery
	}

	if parsedURL.Fragment != "" {
		normalized += "?" + parsedURL.Fragment
	}

	return strings.ToLower(normalized), nil
}
