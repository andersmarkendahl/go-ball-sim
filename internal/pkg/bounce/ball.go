package bounce

import (
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
)

// Ball consist of an Object and image representation
type Ball struct {
	Obj   *objects.Object
	Image *ebiten.Image
}

// New constructor for Ball struct
func New(x, y, vx, vy float64, img *ebiten.Image) (*Ball, error) {
	o, err := objects.New(x, y, vx, vy)
	if err != nil {
		return nil, err
	}
	b := Ball{Obj: o, Image: img}
	return &b, nil
}

// Bounce checks if ball should bounce (invert direction)
func Bounce(b *Ball) error {

	var factor = 0.8

	if b.Obj.Y > float64(ScreenHeight-100) {
		if b.Obj.VY > 0 {
			b.Obj.VY = -b.Obj.VY * factor
		}
	}

	if b.Obj.X > float64(ScreenWidth-100) {
		if b.Obj.VX > 0 {
			b.Obj.VX = -b.Obj.VX * factor
		}
	}

	if b.Obj.X < 100 {
		if b.Obj.VX < 0 {
			b.Obj.VX = -b.Obj.VX * factor
		}
	}
	return nil
}
