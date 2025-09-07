package dto

func cloneMap(original map[string]any) map[string]any {
	cloned := make(map[string]any) // Create a new map
	for k, v := range original {
		k = transformKey(k)
		cloned[k] = v // Copy each key-value pair
	}
	return cloned
}

func deepCloneMap(original map[string]any) map[string]any {
	cloned := make(map[string]any)
	for k, v := range original {
		k = transformKey(k)
		switch v := v.(type) {
		case map[string]any:
			cloned[k] = deepCloneMap(v) // Recursively clone nested maps
		default:
			cloned[k] = v // Copy primitive values
		}
	}
	return cloned
}
