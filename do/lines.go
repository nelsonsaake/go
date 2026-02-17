package do

import "strings"

type LineFlag uint32

const (
	WithoutEmptyLines LineFlag = 1 << iota
)

// splitIntoLines: split string into lines
func splitIntoLines(v string) []string {

	return strings.Split(v, "\n")
}

// removeEmptyLines: remove empty lines and spaces around the none empty lines
func removeEmptyLines(ls []string) []string {

	rs := []string{}

	for _, ln := range ls {
		ln := strings.TrimSpace(ln)
		if len(ln) != 0 {
			rs = append(rs, ln)
		}
	}

	return rs
}

// Lines: split string into lines, and remove all whitespaces around strings, and remove empty lines
func Lines(v string, options ...LineFlag) []string {

	var filter LineFlag = 0
	for _, v := range options {
		filter |= v
	}

	lines := splitIntoLines(v)

	if filter&WithoutEmptyLines != 0 {
		lines = removeEmptyLines(lines)
	}

	return lines
}
