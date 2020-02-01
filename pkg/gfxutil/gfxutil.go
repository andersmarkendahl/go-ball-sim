package gfxutil

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"image"
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

// LoadPngSlice opens a byte slice png image and stores a decoded image
func LoadPngSlice(bslice []byte) (*ebiten.Image, error) {

	img, _, err := image.Decode(bytes.NewReader(bslice))
	if err != nil {
		return nil, err
	}
	image, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	return image, nil

}

// PrintImage draws the image to screen including scale and translation
func PrintImage(screen, img *ebiten.Image, tx, ty, sx, sy float64) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(sx, sy)
	op.GeoM.Translate(tx, ty)
	screen.DrawImage(img, op)

}
