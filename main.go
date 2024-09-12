package main

import "fmt"

func main() {
	normalizedURL, err := normalizeURL("https://blog.boot.dev/path%20with%20spaces")

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Normalized URL:", normalizedURL)
	}
}
