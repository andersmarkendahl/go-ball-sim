package objects

import (
	"errors"
	"github.com/atedja/go-vector"
)

// Object consist of position, velocity
type Object struct {
	X vector.Vector
	V vector.Vector
}

// Position updates position of an object based on current velocity
func (o *Object) Position(dt float64) error {

	if dt <= 0 {
		return errors.New("Timestep (dt) negative or zero")
	}
	tmp := o.V.Clone()
	tmp.Scale(1 / dt)
	o.X = vector.Add(o.X, tmp)
	return nil
}

// Velocity updates velocity of an object based on provided acceleration
func (o *Object) Velocity(ax, ay, dt float64) error {

	if dt <= 0 {
		return errors.New("Timestep (dt) negative or zero")
	}

	tmp := vector.NewWithValues([]float64{ax, ay})
	tmp.Scale(1 / dt)
	o.V = vector.Add(o.V, tmp)
	return nil
}

// New is a constructur of type Object
func New(x, y, vx, vy float64) (*Object, error) {

	X := vector.NewWithValues([]float64{x, y})
	V := vector.NewWithValues([]float64{vx, vy})
	o := Object{X: X, V: V}

	return &o, nil
}
