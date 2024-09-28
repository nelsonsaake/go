package xls2

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/xuri/excelize/v2"
)

// IsRange checks if the provided string is a valid Excel range.
func IsRange(s string) bool {
	// Regular expression to match valid Excel ranges (e.g., A1, B2, A1:B2).
	regex := `(?i)^[A-Z]+\d+(:[A-Z]+\d+)?$`
	matched, err := regexp.MatchString(regex, s)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return matched
}

// IsCell checks if the provided string is a valid Excel cell (case-insensitive).
func IsCell(s string) bool {
	// Regular expression to match a valid Excel cell (e.g., A1, B2, AA10) and ignore case.
	regex := `(?i)^[A-Z]+\d+$`
	matched, err := regexp.MatchString(regex, s)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return matched
}

// SplitRange splits a valid range like "A1:B2" into two cells: "A1" and "B2".
func SplitRange(r string) (string, string, error) {
	if !IsRange(r) {
		return "", "", fmt.Errorf("invalid Excel range: %s", r)
	}

	// Split the range by the colon (":")
	cells := strings.Split(r, ":")
	if len(cells) != 2 {
		return "", "", fmt.Errorf("unexpected range format: %s", r)
	}

	return cells[0], cells[1], nil
}

func minmax(a, b int) (int, int) {
	a64, b64 := float64(a), float64(b)
	a64, b64 = math.Min(a64, b64), math.Max(a64, b64)
	return int(a64), int(b64)
}

// Cells: get __cells in range if it is a range
func __cells(cell1, cell2 string) ([]string, error) {

	cells := []string{}

	c1, r1, err := excelize.CellNameToCoordinates(cell1)
	if err != nil {
		return nil, err
	}

	c2, r2, err := excelize.CellNameToCoordinates(cell2)
	if err != nil {
		return nil, err
	}

	cmin, cmax := minmax(c1, c2)
	rmin, rmax := minmax(r1, r2)

	for i := cmin; i <= cmax; i++ {
		for j := rmin; j <= rmax; j++ {

			colName := ""
			colName, err = excelize.ColumnNumberToName(i)
			if err != nil {
				return nil, err
			}

			cell, err := excelize.JoinCellName(colName, j)

			if err != nil {
				return nil, err
			}

			cells = append(cells, cell)
		}
	}

	return cells, nil
}

// Cells: get cells in range if it is a range or
// returns loc in a string array if it is a cell
// if it two and both are both valid cells,
// it will treat the two as if it was a range
// if it is more than one cell, it will only return an array of valid cells
func Cells(locs ...string) (cells []string, err error) {

	switch len(locs) {

	case 0:
		return nil, errors.New("no location provided")

	case 1:
		loc := locs[0]
		if IsRange(loc) {
			start, end, _ := SplitRange(loc)
			return Cells(start, end)
		}

		if IsCell(loc) {
			return append(cells, loc), nil
		}

		return nil, errors.New("location is invalid")

	case 2:
		start, end := locs[0], locs[1]
		if IsCell(start) && IsCell(end) {
			return __cells(start, end)
		}

		fallthrough

	default:
		for _, loc := range locs {
			if IsCell(loc) {
				cells = append(cells, loc)
			}
		}
	}

	return
}
