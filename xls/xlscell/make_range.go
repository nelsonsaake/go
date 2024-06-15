package xlscell

import (
	"errors"
	"fmt"
)

// MakeRange: makes range from two cells
func MakeRange(start, end string) (string, error) {

	if !IsCell(start) {
		return "", errors.New("start cell is not a valid cell name")
	}

	if !IsCell(end) {
		return "", errors.New("end cell is not a valid cell name")
	}

	return fmt.Sprintf("%s:%s", start, end), nil
}
