package ball

import (
	"github.com/Aoana/ball-sim-go/pkg/gfxutil"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
)

// Ball consist of an Object and image representation
type Ball struct {
	Obj           *objects.Object
	Image         *ebiten.Image
	Radius, Scale float64
}

// New constructor for Ball struct
func New(x, y, vx, vy, scale float64, img *ebiten.Image) (*Ball, error) {
	o, err := objects.New(x, y, vx, vy)
	if err != nil {
		return nil, err
	}
	rect := img.Bounds()
	r := float64(rect.Max.X-rect.Min.X) * scale / 2

	ba := Ball{Obj: o, Image: img, Radius: r, Scale: scale}
	return &ba, nil
}

// Print a ball taking the radius into account
func Print(screen *ebiten.Image, b *Ball) error {
	gfxutil.PrintImage(screen, b.Image, b.Obj.X[0]-b.Radius, b.Obj.X[1]-b.Radius, b.Scale, b.Scale)
	return nil
}

// Boundary checks if ball should bounce within a rectangle (invert direction)
func Boundary(b *Ball, minx, maxx, miny, maxy, factor float64) error {

	if b.Obj.X[0] < minx && b.Obj.V[0] < 0 {
		b.Obj.V[0] = -b.Obj.V[0] * factor
	}
	if b.Obj.X[0] > maxx && b.Obj.V[0] > 0 {
		b.Obj.V[0] = -b.Obj.V[0] * factor
	}
	if b.Obj.X[1] < miny && b.Obj.V[1] < 0 {
		b.Obj.V[1] = -b.Obj.V[1] * factor
	}
	if b.Obj.X[1] > maxy && b.Obj.V[1] > 0 {
		b.Obj.V[1] = -b.Obj.V[1] * factor
	}

	return nil
}
