package axios

import "strings"

// mix replaces placeholders in s using values from env.
// Supported placeholders: {{key}} and {key}
func mix(s string, env map[string]string) string {
	if env == nil {
		return s
	}
	for k, v := range env {
		s = strings.ReplaceAll(s, "{{"+k+"}}", v)
		s = strings.ReplaceAll(s, "{"+k+"}", v)
	}
	return s
}
