package str

import "strconv"

func IsInt(v string) bool {
	_, err := strconv.ParseInt(v, 10, 64)
	return err == nil
}
