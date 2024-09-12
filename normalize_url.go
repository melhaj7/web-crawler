package main

import (
	"errors"
	"net/url"
	"strings"

	"golang.org/x/net/html"
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
		normalized += "#" + parsedURL.Fragment
	}

	return strings.ToLower(normalized), nil
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	base, err := url.Parse(rawBaseURL)
	if err != nil || !base.IsAbs() {
		return nil, errors.New("invalid base url")
	}

	var urls []string
	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href, err := base.Parse(attr.Val)
					if err == nil {
						urls = append(urls, href.String())
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return urls, nil

}
