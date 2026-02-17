package ufs

import "strings"

type ReadLineFlag uint32

const (
	WithoutEmptyLines ReadLineFlag = 1 << iota
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
func lines(v string, options ...ReadLineFlag) []string {

	var filter ReadLineFlag = 0
	for _, v := range options {
		filter |= v
	}

	lines := splitIntoLines(v)

	if filter&WithoutEmptyLines != 0 {
		lines = removeEmptyLines(lines)
	}

	return lines
}

func ReadLines(filepath string, options ...ReadLineFlag) ([]string, error) {

	content, err := ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return lines(content, options...), nil
}
