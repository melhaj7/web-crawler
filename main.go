package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	fmt.Printf("starting crawl of: %s\n", baseURL)

	url := "https://wikipedia.org"
	html, err := getHTML(url)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(html)
	}
}
