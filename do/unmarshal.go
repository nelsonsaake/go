package do

import (
	"encoding/json"
)

// Umarshal: easy way to umarshal http response body,
// be it Response or ResponseRecorder
func Unmarshal(dst any, resp any) error {

	b := Payload(resp)

	if len(b) != 0 {
		return nil
	}

	return json.Unmarshal(b, dst)
}
