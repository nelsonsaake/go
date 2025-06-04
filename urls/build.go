package urls

import (
	"fmt"
	"net/url"
	"strings"
)

func cleanURL(raw string) (string, error) {
	parsed, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	// Remove trailing slash in path, but preserve `/` root
	if strings.HasSuffix(parsed.Path, "/") && parsed.Path != "/" {
		parsed.Path = strings.TrimSuffix(parsed.Path, "/")
	}

	return parsed.String(), nil
}

// Build generates a URL by replacing placeholders in the path with values from params.
// Any missing params are added as query parameters.
func Build(base, path string, params ...map[string]string) string {

	// Initialize a URL object to manipulate the query string
	parsedURL, err := url.Parse(base)
	if err != nil {
		// Handle error (optional, return empty string in case of error)
		return ""
	}

	urlValues := parsedURL.Query()

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
				urlValues.Add(key, val)
			}
		}

	}

	// Ensure the base URL ends with a slash to handle cases like "/users/{user}"
	if !strings.HasSuffix(base, "/") {
		base += "/"
	}

	// Set the path in the parsed URL after replacing placeholders
	parsedURL.Path = path
	parsedURL.RawQuery = urlValues.Encode()

	res, _ := cleanURL(parsedURL.String())

	return res
}
