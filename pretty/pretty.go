package pretty

import (
	"encoding/json"
	"errors"
	"fmt"
)

func JSON(data any) string {

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return ""
	}

	return string(jsonBytes)
}

func Error(data any) error {
	return errors.New(JSON(data))
}

func Print(data any) {
	fmt.Println(JSON(data))
}
