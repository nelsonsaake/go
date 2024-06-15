package xlscell

// IsCell: if it loc is a valid cell name
func IsCell(loc string) bool {
	_, _, err := SplitName(loc)
	return err == nil
}
