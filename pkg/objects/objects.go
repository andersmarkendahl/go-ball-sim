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

// ElasticCollision updates velocity of two objects
// Objects are assumed to have equal mass
func ElasticCollision(o1, o2 *Object) error {

	if o1 == nil || o2 == nil {
		return errors.New("ElasticCollision called with nil object(s)")
	}

	d := vector.Subtract(o1.X, o2.X).Magnitude()
	// Skip collision if objects at same place (divide by zero)
	if d == 0 {
		return nil
	}
	d2 := d * d

	K1, err := vector.Dot(vector.Subtract(o1.V, o2.V), vector.Subtract(o1.X, o2.X))
	// Dot product calculation failed, return error
	if err != nil {
		return err
	}
	K2, err := vector.Dot(vector.Subtract(o2.V, o1.V), vector.Subtract(o2.X, o1.X))
	// Dot product calculation failed, return error
	if err != nil {
		return err
	}

	K1 /= d2
	K2 /= d2

	tmp1 := vector.Subtract(o1.X, o2.X)
	tmp1.Scale(K1)

	tmp2 := vector.Subtract(o2.X, o1.X)
	tmp2.Scale(K2)

	o1.V = vector.Subtract(o1.V, tmp1)
	o2.V = vector.Subtract(o2.V, tmp2)

	return nil

}
