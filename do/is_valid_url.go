package do

import "net/url"

// IsValidURL checks if a string is a valid URL
func IsValidURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}
