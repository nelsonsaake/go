package is

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

// Subject is an interface representing an object that can return values by key
type Subject interface {
	Get(key string) string
}

var (
	toMapCache sync.Map // map[uintptr]map[string]any
)

// Entry point
func HasProperty(subject Subject, property string, rules map[string][]map[string]any) bool {
	ruleGroup, ok := rules[property]
	if !ok {
		return false
	}
	return MatchesAny(subject, ruleGroup)
}

// Match any rule in a rule group
func MatchesAny(subject Subject, rules []map[string]any) bool {
	for _, rule := range rules {
		if matchesAll(subject, rule) {
			return true
		}
	}
	return false
}

// Match one rule (all entries must match)
func matchesAll(subject Subject, rule map[string]any) bool {
	for key, value := range rule {
		if !matches(subject, key, value) {
			return false
		}
	}
	return true
}

// Match a single rule entry
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

// Match exact value or array of values (case-insensitive)
func matchesValue(value any, fieldValue string) bool {
	fieldValue = strings.ToLower(fieldValue)
	switch v := value.(type) {
	case string:
		return strings.ToLower(v) == fieldValue
	case []any:
		for _, item := range v {
			if strItem, ok := item.(string); ok && strings.ToLower(strItem) == fieldValue {
				return true
			}
		}
	}
	return false
}

// Match regex or array of regex patterns (case-insensitive)
func matchesPattern(value any, fieldValue string) bool {
	switch v := value.(type) {
	case string:
		return matchRegex(v, fieldValue)
	case []any:
		for _, item := range v {
			if strItem, ok := item.(string); ok && matchRegex(strItem, fieldValue) {
				return true
			}
		}
	}
	return false
}

// Case-insensitive regex match (adds (?i) only if not already present)
func matchRegex(pattern, value string) bool {

	// Add case-insensitive flag only if not already present
	if !isCaseInsensitive(pattern) {
		pattern = "(?i)" + pattern
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(fmt.Errorf("failed to compile regex: %v", err))
	}

	return re.MatchString(value)
}

// Detects if a pattern already contains a case-insensitive flag
func isCaseInsensitive(pattern string) bool {
	// Match either (?i) at the beginning or (?i:...) group
	return strings.HasPrefix(pattern, "(?i)") || strings.HasPrefix(pattern, "(?i:")
}

// ToMap converts any value into a map[string]any with caching.
func ToMap(value any) map[string]any {
	if value == nil {
		return nil
	}

	// Try cache based on pointer identity
	ptr := reflect.ValueOf(value)
	if ptr.Kind() == reflect.Pointer {
		if cached, ok := toMapCache.Load(ptr.Pointer()); ok {
			return cached.(map[string]any)
		}
	}

	var result map[string]any

	switch v := value.(type) {
	case map[string]any:
		result = v

	case Subject:
		result = make(map[string]any)
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Pointer {
			val = val.Elem()
		}
		typ := val.Type()
		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			if field.IsExported() {
				result[field.Name] = val.Field(i).Interface()
			}
		}

	default:
		var m map[string]any
		b, err := json.Marshal(v)
		if err != nil {
			return nil
		}
		if err := json.Unmarshal(b, &m); err != nil {
			return nil
		}
		result = m
	}

	// Memoize if it's a pointer
	if ptr.Kind() == reflect.Pointer {
		toMapCache.Store(ptr.Pointer(), result)
	}

	return result
}
