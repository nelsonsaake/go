package tests

import (
	"testing"

	"github.com/nelsonsaake/go/zip"
)

func TestZip(t *testing.T) {

	originalfile := "bin/1"
	zippedfile := "bin/zipped/1.zip"
	unzippedfile := "bin/unzipped"

	p, err := zip.Zip(zippedfile, originalfile)
	if err != nil {
		t.Fatalf("Error zipping files: %v", err)
	}

	err = zip.Unzip(p, unzippedfile)
	if err != nil {
		t.Fatalf("Error unzipping files: %v", err)
	}
}
