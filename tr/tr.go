package tr

import (
	"fmt"
	"regexp"
	"strings"
)

// Subject is an interface representing an object that can return values by key
type Subject interface {
	Get(key string) string
}

// MatchesAny checks all rules and returns the value of `target` in the first matching rule.
// It ignores the target field when evaluating matches.
func MatchesAny(subject Subject, target string, rules []map[string]any) any {

	for _, rule := range rules {
		if matchesAll(subject, target, rule) {
			return rule[target]
		}
	}
	return nil
}

// matchesAll returns true if all fields in a rule (except the target key) match
func matchesAll(subject Subject, target string, rule map[string]any) bool {
	for key, value := range rule {
		if key == target {
			continue
		}
		if !matches(subject, key, value) {
			return false
		}
	}
	return true
}

// matches determines if a subject field matches a rule key and value
// Supports "key" and "keyPattern" matching
func matches(subject Subject, key string, value any) bool {
	isPattern := strings.HasSuffix(strings.ToLower(key), "pattern")
	lookupKey := key
	if isPattern {
		lookupKey = key[:len(key)-7]
	}

	fieldVal := subject.Get(lookupKey)
	if isPattern {
		return matchesPattern(value, fieldVal)
	}
	return matchesValue(value, fieldVal)
}

// matchesValue matches exact value(s), case-insensitively
func matchesValue(value any, fieldValue string) bool {
	fieldValue = strings.ToLower(fieldValue)
	switch v := value.(type) {
	case string:
		return strings.ToLower(v) == fieldValue
	case []any:
		for _, item := range v {
			if str, ok := item.(string); ok && strings.ToLower(str) == fieldValue {
				return true
			}
		}
	}
	return false
}

// matchesPattern matches regex pattern(s), case-insensitively
func matchesPattern(value any, fieldValue string) bool {
	switch v := value.(type) {
	case string:
		return matchRegex(v, fieldValue)
	case []any:
		for _, item := range v {
			if str, ok := item.(string); ok && matchRegex(str, fieldValue) {
				return true
			}
		}
	}
	return false
}

// matchRegex adds (?i) if not already present, then compiles and matches
func matchRegex(pattern, value string) bool {
	if !isCaseInsensitive(pattern) {
		pattern = "(?i)" + pattern
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(fmt.Errorf("regex compile failed: %v", err))
	}
	return re.MatchString(value)
}

// isCaseInsensitive checks if pattern already starts with (?i) or (?i:
func isCaseInsensitive(pattern string) bool {
	return strings.HasPrefix(pattern, "(?i)") || strings.HasPrefix(pattern, "(?i:")
}
