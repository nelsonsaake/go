package xlsutil

import "math"

// calcOffset: calculate B(a picture) offset in a A(a cell in a spreadsheet)
//
// calculates where B start in A or the position of B in A.
// if object B is inside A such that their centers are aligned and
// the length of A and B is provided calcOffset caclulates where B starts.
// the length of A or B can be interchanged as parameters
// the results of this calculation is scalar
//
// can be used to calculate B(a picture) offset in a A(a cell in a spreadsheet)
func calcOffset(lb, la float64) float64 {
	return math.Abs((lb - la) / 2.0)
}
