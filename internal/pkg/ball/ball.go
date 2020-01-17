package ball

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

// Boundary checks if ball should bounce within a rectangle (invert direction)
func Boundary(b *Ball, minx, maxx, miny, maxy, factor float64) error {

	if b.Obj.X < minx && b.Obj.VX < 0 {
		b.Obj.VX = -b.Obj.VX * factor
	}
	if b.Obj.X > maxx && b.Obj.VX > 0 {
		b.Obj.VX = -b.Obj.VX * factor
	}
	if b.Obj.Y < miny && b.Obj.VY < 0 {
		b.Obj.VY = -b.Obj.VY * factor
	}
	if b.Obj.Y > maxy && b.Obj.VY > 0 {
		b.Obj.VY = -b.Obj.VY * factor
	}

	return nil
}
