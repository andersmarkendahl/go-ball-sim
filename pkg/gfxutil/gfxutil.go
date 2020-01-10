package gfxutil

import (
	"github.com/hajimehoshi/ebiten"
	"image/png"
	"os"
)

// LoadPng opens a png image and stores a decoded image
func LoadPng(pngimage string) (*ebiten.Image, error) {

	file, err := os.Open(pngimage)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	image, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	return image, nil

}
