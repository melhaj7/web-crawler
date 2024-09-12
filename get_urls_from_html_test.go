package main

import (
	"reflect"
	"testing"
)

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
