package main

import (
	"reflect"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "URL with no scheme",
			inputURL: "blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "URL with http scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "URL with trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path/",
		},
		{
			name:     "URL with subdomain and query params",
			inputURL: "https://sub.blog.boot.dev/path?query=123",
			expected: "sub.blog.boot.dev/path?query=123",
		},
		{
			name:     "URL with port number",
			inputURL: "https://blog.boot.dev:8080/path",
			expected: "blog.boot.dev:8080/path",
		},
		{
			name:     "URL with fragment",
			inputURL: "https://blog.boot.dev/path#section1",
			expected: "blog.boot.dev/path#section1",
		},
		{
			name:     "Empty input URL",
			inputURL: "",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		htmlBody  string
		baseURL   string
		expected  []string
		expectErr bool
	}{
		{
			name: "absolute and relative URLs",
			htmlBody: `
<html>
	<body>
		<a href="/path/one">Link One</a>
		<a href="https://other.com/path/one">Link Two</a>
	</body>
</html>
`,
			baseURL: "https://blog.boot.dev",
			expected: []string{
				"https://blog.boot.dev/path/one",
				"https://other.com/path/one",
			},
		},
		{
			name: "only relative URLs",
			htmlBody: `
<html>
	<body>
		<a href="/path/two">Link Three</a>
	</body>
</html>
`,
			baseURL: "https://example.com",
			expected: []string{
				"https://example.com/path/two",
			},
		},
		{
			name: "invalid base URL",
			htmlBody: `
<html>
	<body>
		<a href="/path/four">Link Four</a>
	</body>
</html>
`,
			baseURL:   "invalid-url",
			expectErr: true,
		},
		// 		{
		// 			name: "no links",
		// 			htmlBody: `
		// <html>
		// 	<body>
		// 		<p>No links here.</p>
		// 	</body>
		// </html>
		// `,
		// 			baseURL:  "https://example.com",
		// 			expected: []string{},
		// 		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.htmlBody, tc.baseURL)
			if (err != nil) != tc.expectErr {
				t.Errorf("Test %v - '%s' FAIL: expected error: %v, got: %v", i, tc.name, tc.expectErr, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected URLs: %v, got: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
