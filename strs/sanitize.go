package strs

import "regexp"

// SanitizeFileName replaces invalid characters in a file name with underscores.
func SanitizeFileName(name string) string {
	// Replace spaces and special characters with underscores

	var invalid = regexp.MustCompile(`[ <>:"/\\|?*]`)
	sanitized := invalid.ReplaceAllString(name, "_")
	return sanitized
}
