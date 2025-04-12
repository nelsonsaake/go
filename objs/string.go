package objs

import "encoding/json"

func String(m map[string]any) string {
	jsonBytes, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}
