package ball

import (
	"errors"
	"github.com/Aoana/ball-sim-go/pkg/gfxutil"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/atedja/go-vector"
	"github.com/hajimehoshi/ebiten"
)

// BallList is the global slice of balls
var BallList []*Ball

// Ball consist of an Object and image representation
type Ball struct {
	Obj           *objects.Object
	Image         *ebiten.Image
	Radius, Scale float64
	Active        bool
}

// New is a constructor for Ball struct
func New(x, y, vx, vy, scale float64, img *ebiten.Image) (*Ball, error) {
	o, err := objects.New(x, y, vx, vy)
	if err != nil {
		return nil, err
	}
	rect := img.Bounds()
	r := float64(rect.Max.X-rect.Min.X) * scale / 2

	ba := Ball{Obj: o, Image: img, Radius: r, Scale: scale, Active: true}
	return &ba, nil
}

// Add adds a ball to global list
func Add(b *Ball) error {

	BallList = append(BallList, b)
	return nil
}

// Remove removes ball from global list
func Remove(index int) error {

	if index < 0 || index >= len(BallList) {
		return errors.New("Out of range index in BallList")
	}

	BallList[index] = BallList[len(BallList)-1]
	BallList[len(BallList)-1] = nil
	BallList = BallList[:len(BallList)-1]

	return nil
}

// Print a ball taking the radius into account
func Print(screen *ebiten.Image, b *Ball) error {
	gfxutil.PrintImage(screen, b.Image, b.Obj.X[0]-b.Radius, b.Obj.X[1]-b.Radius, b.Scale, b.Scale)
	return nil
}

// Boundary checks if ball should bounce within a rectangle (invert direction)
// Calculation based on collision with fixed surface
// Coefficent of restitution set by factor (if 1 elastic collision)
func Boundary(b *Ball, minx, maxx, miny, maxy, factor float64) error {

	if b.Obj.X[0] < (minx+b.Radius) && b.Obj.V[0] < 0 {
		b.Obj.V[0] = -b.Obj.V[0] * factor
	}
	if b.Obj.X[0] > (maxx-b.Radius) && b.Obj.V[0] > 0 {
		b.Obj.V[0] = -b.Obj.V[0] * factor
	}
	if b.Obj.X[1] < (miny+b.Radius) && b.Obj.V[1] < 0 {
		b.Obj.V[1] = -b.Obj.V[1] * factor
	}
	if b.Obj.X[1] > (maxy-b.Radius) && b.Obj.V[1] > 0 {
		b.Obj.V[1] = -b.Obj.V[1] * factor
	}

	return nil
}

// Collide updates balls based on collision to other balls
// Calculation based on elastic collision with equal mass
func Collide(b1, b2 *Ball) {

	d := vector.Subtract(b1.Obj.X, b2.Obj.X).Magnitude()

	if d > b1.Radius+b2.Radius {
		// Balls do not collide
		return
	}

	K1, _ := vector.Dot(vector.Subtract(b1.Obj.V, b2.Obj.V), vector.Subtract(b1.Obj.X, b2.Obj.X))
	K2, _ := vector.Dot(vector.Subtract(b2.Obj.V, b1.Obj.V), vector.Subtract(b2.Obj.X, b1.Obj.X))
	K1 /= (d * d)
	K2 /= (d * d)

	tmp1 := vector.Subtract(b1.Obj.X, b2.Obj.X)
	tmp1.Scale(K1)

	tmp2 := vector.Subtract(b2.Obj.X, b1.Obj.X)
	tmp2.Scale(K2)

	b1.Obj.V = vector.Subtract(b1.Obj.V, tmp1)
	b2.Obj.V = vector.Subtract(b2.Obj.V, tmp2)

}
