package objects

import (
	"errors"
)

// Object consist of position, velocity
type Object struct {
	X, Y   float64
	VX, VY float64
}

// Position updates position of an object based on current velocity
func (o *Object) Position(dt float64) error {

	if dt <= 0 {
		return errors.New("Timestep (dt) negative or zero")
	}

	o.X = o.X + o.VX/dt
	o.Y = o.Y + o.VY/dt
	return nil
}

// Velocity updates velocity of an object based on provided acceleration
func (o *Object) Velocity(ax, ay, dt float64) error {

	if dt <= 0 {
		return errors.New("Timestep (dt) negative or zero")
	}

	o.VX = o.VX + ax/dt
	o.VY = o.VY + ay/dt
	return nil
}

// New is a constructur of type Object
func New(x, y, vx, vy float64) (*Object, error) {

	o := Object{X: x, Y: y, VX: vx, VY: vy}

	return &o, nil
}
