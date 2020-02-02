package objects

import (
	"math"
	"testing"
)

// Round to 3 decimals to ignore 10^-15 diffs
func r(x float64) float64 {
	return math.Round(x*1000) / 1000
}

func isDifferent(a, b Object) bool {

	if r(a.X[0]) == r(b.X[0]) && r(a.X[1]) == r(b.X[1]) && r(a.V[0]) == r(b.V[0]) && r(a.V[1]) == r(b.V[1]) {
		return false
	}
	return true
}

var fostring = "\nin: %+v\nwanted: %+v\nresult: %+v"

func TestNew(t *testing.T) {
	for _, c := range []struct {
		want         Object
		x, y, vx, vy float64
	}{
		// Positive and negative permutations
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, 0, 0, 0, 0},
		{Object{X: []float64{-1, 2}, V: []float64{3, 4}}, -1, 2, 3, 4},
		{Object{X: []float64{1, -2}, V: []float64{3, 4}}, 1, -2, 3, 4},
		{Object{X: []float64{1, 2}, V: []float64{-3, 4}}, 1, 2, -3, 4},
		{Object{X: []float64{1, 2}, V: []float64{3, -4}}, 1, 2, 3, -4},
		{Object{X: []float64{-1, -2}, V: []float64{3, 4}}, -1, -2, 3, 4},
		{Object{X: []float64{1, -2}, V: []float64{-3, 4}}, 1, -2, -3, 4},
		{Object{X: []float64{1, 2}, V: []float64{-3, -4}}, 1, 2, -3, -4},
		{Object{X: []float64{-1, 2}, V: []float64{3, -4}}, -1, 2, 3, -4},
		// Float values
		{Object{X: []float64{1.1, 2.2}, V: []float64{3.3, 4.4}}, 1.1, 2.2, 3.3, 4.4},
		{Object{X: []float64{-1.1, 2.2}, V: []float64{3.3, 4.4}}, -1.1, 2.2, 3.3, 4.4},
		{Object{X: []float64{1.1, -2.2}, V: []float64{3.3, 4.4}}, 1.1, -2.2, 3.3, 4.4},
		{Object{X: []float64{1.1, 2.2}, V: []float64{-3.3, 4.4}}, 1.1, 2.2, -3.3, 4.4},
		{Object{X: []float64{1.1, 2.2}, V: []float64{3.3, -4.4}}, 1.1, 2.2, 3.3, -4.4},
	} {
		got, err := New(c.x, c.y, c.vx, c.vy)
		if err != nil {
			t.Errorf("New(%f, %f, %f, %f) %s", c.x, c.y, c.vx, c.vy, err)
		}
		if isDifferent(*got, c.want) {
			t.Errorf("New(%f, %f, %f, %f) \nwanted: %+v\nresult: %+v", c.x, c.y, c.vx, c.vy, c.want, *got)
		}
	}

}

func TestPosition(t *testing.T) {
	// Positive tests
	for _, c := range []struct {
		in, want Object
		dt       float64
	}{
		// Position update in all directions
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{0, 0}, V: []float64{0, 0}}, 1},
		{Object{X: []float64{0, 0}, V: []float64{1, 0}}, Object{X: []float64{1, 0}, V: []float64{1, 0}}, 1},
		{Object{X: []float64{0, 0}, V: []float64{0, 1}}, Object{X: []float64{0, 1}, V: []float64{0, 1}}, 1},
		{Object{X: []float64{0, 0}, V: []float64{1, 1}}, Object{X: []float64{1, 1}, V: []float64{1, 1}}, 1},
		{Object{X: []float64{0, 0}, V: []float64{-1, 1}}, Object{X: []float64{-1, 1}, V: []float64{-1, 1}}, 1},
		{Object{X: []float64{0, 0}, V: []float64{1, -1}}, Object{X: []float64{1, -1}, V: []float64{1, -1}}, 1},
		{Object{X: []float64{0, 0}, V: []float64{-1, -1}}, Object{X: []float64{-1, -1}, V: []float64{-1, -1}}, 1},
		// Float values
		{Object{X: []float64{1.1, 2.2}, V: []float64{1.1, 2.2}}, Object{X: []float64{2.2, 4.4}, V: []float64{1.1, 2.2}}, 1},
		{Object{X: []float64{1.1, 2.2}, V: []float64{-1.1, 2.2}}, Object{X: []float64{0, 4.4}, V: []float64{-1.1, 2.2}}, 1},
		{Object{X: []float64{1.1, 2.2}, V: []float64{1.1, -2.2}}, Object{X: []float64{2.2, 0}, V: []float64{1.1, -2.2}}, 1},
		{Object{X: []float64{1.1, 2.2}, V: []float64{-1.1, -2.2}}, Object{X: []float64{0, 0}, V: []float64{-1.1, -2.2}}, 1},
		// dt variation
		{Object{X: []float64{100, 200}, V: []float64{20, 40}}, Object{X: []float64{102, 204}, V: []float64{20, 40}}, 10},
		{Object{X: []float64{100, 200}, V: []float64{20, 40}}, Object{X: []float64{108, 216}, V: []float64{20, 40}}, 2.5},
		{Object{X: []float64{100, 200}, V: []float64{20, 40}}, Object{X: []float64{112.5, 225}, V: []float64{20, 40}}, 1.6},
	} {
		orig := c.in
		c.in.Position(c.dt)
		if isDifferent(c.in, c.want) {
			t.Errorf("Position(%f)"+fostring, c.dt, orig, c.want, c.in)
		}
	}
	// Negative tests
	for _, c := range []struct {
		in Object
		dt float64
	}{
		// Timestep cannot be zero or negative
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, 0},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, -1},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, -40000},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, -3.1416},
	} {
		err := c.in.Position(c.dt)
		if err == nil {
			t.Errorf("Position(%f) should fail", c.dt)
		}
	}
}

func TestVelocity(t *testing.T) {
	for _, c := range []struct {
		in, want   Object
		ax, ay, dt float64
	}{
		// Velocity update in all directions
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{0, 0}, V: []float64{0, 0}}, 0, 0, 1},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{0, 0}, V: []float64{1, 0}}, 1, 0, 1},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{0, 0}, V: []float64{0, 1}}, 0, 1, 1},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{0, 0}, V: []float64{1, 1}}, 1, 1, 1},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{0, 0}, V: []float64{-1, 0}}, -1, 0, 1},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{0, 0}, V: []float64{0, -1}}, 0, -1, 1},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{0, 0}, V: []float64{-1, -1}}, -1, -1, 1},
		// Float values
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{0, 0}, V: []float64{1.5, 2.5}}, 1.5, 2.5, 1},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{0, 0}, V: []float64{-1.5, 2.5}}, -1.5, 2.5, 1},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{0, 0}, V: []float64{1.5, -2.5}}, 1.5, -2.5, 1},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{0, 0}, V: []float64{-1.5, -2.5}}, -1.5, -2.5, 1},
		// dt variation
		{Object{X: []float64{100, 200}, V: []float64{20, 40}}, Object{X: []float64{100, 200}, V: []float64{21, 41}}, 10, 10, 10},
		{Object{X: []float64{100, 200}, V: []float64{20, 40}}, Object{X: []float64{100, 200}, V: []float64{24, 44}}, 10, 10, 2.5},
		{Object{X: []float64{100, 200}, V: []float64{20, 40}}, Object{X: []float64{100, 200}, V: []float64{26.25, 46.25}}, 10, 10, 1.6},
	} {
		orig := c.in
		c.in.Velocity(c.ax, c.ay, c.dt)
		if isDifferent(c.in, c.want) {
			t.Errorf("Velocity(%f, %f, %f)"+fostring, c.ax, c.ay, c.dt, orig, c.want, c.in)
		}
	}
	// Argument tests (negative tests)
	for _, c := range []struct {
		in         Object
		ax, ay, dt float64
	}{
		// Timestep cannot be zero or negative
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, 0, 0, 0},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, 0, 0, -1},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, 0, 0, -40000},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, 0, 0, -3.1416},
	} {
		err := c.in.Velocity(c.ax, c.ay, c.dt)
		if err == nil {
			t.Errorf("Velocity(%f, %f, %f) should fail", c.ax, c.ay, c.dt)
		}
	}

}

func TestElasticCollision(t *testing.T) {
	for _, c := range []struct {
		in1, in2     Object
		want1, want2 Object
	}{
		// No moving Objects
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{1, 1}, V: []float64{0, 0}},
			Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{1, 1}, V: []float64{0, 0}}},
		// Objects in same place (no change)
		{Object{X: []float64{5, 5}, V: []float64{100, 100}}, Object{X: []float64{5, 5}, V: []float64{-2, 2}},
			Object{X: []float64{5, 5}, V: []float64{100, 100}}, Object{X: []float64{5, 5}, V: []float64{-2, 2}}},
		// One Object at rest
		{Object{X: []float64{0, 0}, V: []float64{1.5, 0}}, Object{X: []float64{1, 0}, V: []float64{0, 0}},
			Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{1, 0}, V: []float64{1.5, 0}}},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{1, 0}, V: []float64{-0.9, 0}},
			Object{X: []float64{0, 0}, V: []float64{-0.9, 0}}, Object{X: []float64{1, 0}, V: []float64{0, 0}}},
		{Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{0, 1}, V: []float64{0, -1000.55}},
			Object{X: []float64{0, 0}, V: []float64{0, -1000.55}}, Object{X: []float64{0, 1}, V: []float64{0, 0}}},
		{Object{X: []float64{0, 0}, V: []float64{1, 1}}, Object{X: []float64{1, 1}, V: []float64{0, 0}},
			Object{X: []float64{0, 0}, V: []float64{0, 0}}, Object{X: []float64{1, 1}, V: []float64{1, 1}}},
		// Frontal collision
		{Object{X: []float64{0, 0}, V: []float64{1, 0}}, Object{X: []float64{1, 0}, V: []float64{-1, 0}},
			Object{X: []float64{0, 0}, V: []float64{-1, 0}}, Object{X: []float64{1, 0}, V: []float64{1, 0}}},
		{Object{X: []float64{0, 0}, V: []float64{100, 100}}, Object{X: []float64{1, 1}, V: []float64{-100, -100}},
			Object{X: []float64{0, 0}, V: []float64{-100, -100}}, Object{X: []float64{1, 1}, V: []float64{100, 100}}},
	} {
		orig1 := c.in1
		orig2 := c.in2
		err := ElasticCollision(&c.in1, &c.in2)
		if err != nil {
			t.Errorf("ElasticCollision(%+vf, %+v) error %s", c.in1, c.in2, err)
		}
		if isDifferent(c.in1, c.want1) {
			t.Errorf("ElasticCollision(%+vf, %+v) %+vf != %+v", orig1, orig2, c.in1, c.want1)
		}
		if isDifferent(c.in2, c.want2) {
			t.Errorf("ElasticCollision(%+vf, %+v) %+vf != %+v", orig1, orig2, c.in2, c.want2)
		}
	}

}
