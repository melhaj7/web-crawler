package main

import (
	"fmt"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return fmt.Sprintf("HTTP status: %v", resp.StatusCode), nil
}
