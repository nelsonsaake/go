package xlsutil

import "encoding/json"

type formatSheet map[string]any

func NewFormatSheet() formatSheet {
	return map[string]any{}
}

func (f formatSheet) AsString() string {
	bs, err := json.MarshalIndent(f, "", "\t")
	if err != nil {
		return ""
	}
	return string(bs)
}

func (f formatSheet) Set(key string, value any) {
	f[key] = value
}

func (f formatSheet) IsSet(key string) bool {
	_, ok := f[key]
	return ok
}

func (f formatSheet) Remove(key string) {
	delete(f, key)
}
