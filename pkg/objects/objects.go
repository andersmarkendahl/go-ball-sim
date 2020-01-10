package objects

import (
	"github.com/hajimehoshi/ebiten"
)

// Object struct
// Consist of position, velocity and image representation
type Object struct {
	X, Y   float64
	VX, VY float64
	Image  *ebiten.Image
}

// Position updates position of an object based on current velocity
func (o *Object) Position(dt float64) error {
	o.X = o.X + o.VX/dt
	o.Y = o.Y + o.VY/dt
	return nil
}

// Velocity updates velocity of an object based on provided acceleration
func (o *Object) Velocity(a, dt float64) error {
	o.VY = o.VY + a/dt
	return nil
}

// New is a constructur of type Object
func New(x, y, vx, vy float64, image *ebiten.Image) (*Object, error) {

	o := Object{X: x, Y: y, VX: vx, VY: vy, Image: image}

	return &o, nil
}
