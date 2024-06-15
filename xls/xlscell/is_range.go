package xlscell

// IsRange: check if loc is a range
func IsRange(loc string) bool {
	_, _, err := SplitRange(loc)
	return err == nil
}
