package xlsutil

const (
	PhotoWidth  = 150
	PhotoHeight = 200
)

// calcScale: calculate the scale factor given the target dimension and the actual dimension
// this means if the scale is multiplied by the actual dimension, it should result in the target dimension
func calcScale(targetDimension, actualDimension int) float64 {
	return float64(targetDimension) / float64(actualDimension)
}

// calcScaleWD: calculate scale width default,
// for calculating scale factor using PhotoWidth as the target dimension
// and the provided value as the actual dimension
func calcScaleWD(w int) float64 {
	return calcScale(PhotoWidth, w)
}

// calcScaleWD: calculate scale height default,
// for calculating scale factor using PhotoHeight as the target dimension
// and the provided value as the actual dimension
func calcScaleHD(h int) float64 {
	return calcScale(PhotoHeight, h)
}
