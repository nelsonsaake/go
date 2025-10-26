package repo

import "encoding/json"

// toMap converts any Go data structure to a map[string]any
// by marshaling to JSON and then unmarshaling into a map.
func toMap(data any) (map[string]any, error) {

	res := make(map[string]any)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(jsonData, &res)

	return res, nil
}
