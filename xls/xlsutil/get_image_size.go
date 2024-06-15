package xlsutil

import (
	"image"
	"io"
	"os"

	"golang.org/x/image/webp"
)

func getImageSize(reader io.Reader) (w, h int, err error) {

	imgConfig, _, err := image.DecodeConfig(reader)
	if err == nil {
		return imgConfig.Width, imgConfig.Height, nil
	}

	imgConfig, err = webp.DecodeConfig(reader)
	if err == nil {
		return imgConfig.Width, imgConfig.Height, nil
	}

	return 0, 0, err
}

func getImageSizeFromFile(imgfile string) (w, h int, err error) {

	reader, err := os.Open(imgfile)
	if err != nil {
		return
	}

	return getImageSize(reader)
}
