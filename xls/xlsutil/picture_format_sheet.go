package xlsutil

import (
	"io"
)

type pictureFormatSheet struct {
	formatSheet

	// picture width
	w int

	// picture height
	h int
}

func NewPictureFormatSheet(reader io.Reader) (*pictureFormatSheet, error) {

	w, h, err := getImageSize(reader)
	if err != nil {
		return nil, err
	}

	p := &pictureFormatSheet{
		formatSheet: NewFormatSheet(),
		w:           w,
		h:           h,
	}

	return p, nil
}

func (p *pictureFormatSheet) Scale() *pictureFormatSheet {
	p.Set("x_scale", calcScaleWD(p.w))
	p.Set("y_scale", calcScaleHD(p.h))

	return p
}

func (p *pictureFormatSheet) Offset(y int) *pictureFormatSheet {
	p.Set("x_offset", 10)
	p.Set("y_offset", y)

	return p
}

func (p *pictureFormatSheet) Fit() *pictureFormatSheet {
	p.Set("autofit", true)
	return p
}
