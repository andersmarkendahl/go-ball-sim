package gfxutil

import (
	"testing"
)

// LoadPng opens a png image and stores a decoded image
func TestLoadPng(t *testing.T) {

	// Positive tests
	for _, c := range []string{
		"testdata/example.png",
	} {
		img, err := LoadPng(c)
		if err != nil {
			t.Errorf("LoadPng(%s): %s", c, err)
		}
		if img == nil {
			t.Errorf("LoadPng(%s): empty image", c)
		}
	}
	// Negative tests
	for _, c := range []string{
		"testdata/donotexist.png",
		"testdata/example.jpg",
		"testdata/garbage",
	} {
		_, err := LoadPng(c)
		if err == nil {
			t.Errorf("LoadPng(%s): loadable", c)
		}
	}
}

func TestPrintImage(t *testing.T) {

	// Nothing to test, ebiten does not provide error codes here

}
