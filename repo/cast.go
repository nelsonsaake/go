package repo

import "encoding/json"

// cast converts any Go data structure to a pointer of type T
// by marshaling to JSON and then unmarshaling into T.
func cast[T any](data any) (*T, error) {

	res := new(T)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
