package urls

import (
	"fmt"
	"net/url"
	"strings"
)

// Build generates a URL by replacing placeholders in the path with values from params.
// Any missing params are added as query parameters.
func Build(base, path string, params ...map[string]string) string {
	// Initialize a URL object to manipulate the query string
	parsedURL, err := url.Parse(base)
	if err != nil {
		// Handle error (optional, return empty string in case of error)
		return ""
	}

	// Iterate over all provided params
	for _, params := range params {
		// apply parameters to url
		for key, val := range params {
			// placeholder
			placeholder := fmt.Sprintf("{%s}", key)

			// check if the key is path param
			if strings.Contains(path, placeholder) {
				// if yes, set the value
				path = strings.ReplaceAll(path, placeholder, val)
			} else {
				// otherwise, add it as a query param
				parsedURL.Query().Add(key, val)
			}
		}

	}

	// Ensure the base URL ends with a slash to handle cases like "/users/{user}"
	if !strings.HasSuffix(base, "/") {
		base += "/"
	}

	// Set the path in the parsed URL after replacing placeholders
	parsedURL.Path = path
	parsedURL.RawQuery = parsedURL.Query().Encode()

	return parsedURL.String()
}
