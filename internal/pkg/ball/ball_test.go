package ball

import (
	"github.com/Aoana/go-ball-sim/assets/images"
	"github.com/Aoana/go-ball-sim/pkg/gfxutil"
	"github.com/hajimehoshi/ebiten"
	"math"
	"testing"
)

var img *ebiten.Image

// Round to 3 decimals to ignore 10^-15 diffs
func r(x float64) float64 {
	return math.Round(x*1000) / 1000
}

// Helper function to ball vs values
func isDifferent(x, y, vx, vy float64, b *Ball) bool {

	if x == r(b.Obj.X[0]) && y == r(b.Obj.X[1]) && vx == r(b.Obj.V[0]) && vy == r(b.Obj.V[1]) {
		return false
	}
	return true
}

// Helper function to check if ball is present in BallList
func isPresent(x, y, vx, vy, scale float64) bool {
	for i := range BallList {
		b := BallList[i]
		if x == b.Obj.X[0] && y == b.Obj.X[1] && vx == b.Obj.V[0] && vy == b.Obj.V[1] && scale == b.Scale {
			return true
		}
	}
	return false
}

// Loads a test image by default
func init() {
	img, _ = gfxutil.LoadPngSlice(images.ImageStar)
}

// Test function to create, add and remove balls to the global list
func TestList(t *testing.T) {

	for _, c := range []struct {
		x, y, vx, vy, scale float64
	}{
		// Example ball starting values
		{1.1, 2.2, 3.3, 4.4, 0.5},
		{-1.1, 50, 50000, 0, 2.323},
		{0, 0, 0, 0, -99.9},
	} {
		b, err := New(c.x, c.y, c.vx, c.vy, c.scale, img)
		if err != nil {
			t.Errorf("New(%f, %f, %f, %f, %f, %p) error %s", c.x, c.y, c.vx, c.vy, c.scale, img, err)
		}
		err = Add(b)
		if err != nil {
			t.Errorf("Add(%+v) error %s", b, err)
		}
		if !isPresent(c.x, c.y, c.vx, c.vy, c.scale) {
			t.Errorf("Ball missing (%f, %f, %f, %f, %f, %p)", c.x, c.y, c.vx, c.vy, c.scale, img)
		}
	}
	err := Remove(1)
	if err != nil {
		t.Error("Remove(1) failed ", err)
	}
	x, y, vx, vy, scale := -1.1, 50.0, 50000.0, 0.0, 2.323
	if isPresent(x, y, vx, vy, scale) {
		t.Errorf("Ball present (%f, %f, %f, %f, %f, %p)", x, y, vx, vy, scale, img)
	}
}

// Test function to create, add and remove balls to the global list
func TestBoundary(t *testing.T) {

	for _, c := range []struct {
		x0, y0, vx0, vy0 float64
		x, y, vx, vy     float64
		factor           float64
	}{
		// In corners direction is inwards
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 99, 1, -1, 1, 99, 1, -1, 1},
		{99, 1, -1, 1, 99, 1, -1, 1, 1},
		{99, 99, -1, -1, 99, 99, -1, -1, 1},
		// In corners and bounce
		{1, 1, -1, -1, 1, 1, 1, 1, 1},
		{1, 99, -1, 1, 1, 99, 1, -1, 1},
		{99, 1, 1, -1, 99, 1, -1, 1, 1},
		{99, 99, 1, 1, 99, 99, -1, -1, 1},
		// Factor variations
		{1, 1, 1, 1, 1, 1, 1, 1, 0.5},
		{1, 1, 1, 1, 1, 1, 1, 1, 10},
		{1, 1, -1, -1, 1, 1, 0.5, 0.5, 0.5},
		{1, 1, -1, -1, 1, 1, 10, 10, 10},
	} {
		b, err := New(c.x0, c.y0, c.vx0, c.vy0, 1.0, img)
		if err != nil {
			t.Errorf("New(%f, %f, %f, %f, %f, %p) error %s", c.x0, c.y0, c.vx0, c.vy0, 1.0, img, err)
		}
		err = Boundary(b, 0.0, 100.0, 0.0, 100.0, c.factor)
		if err != nil {
			t.Errorf("Boundary(%+v, 0, 1, 0, 1) error %s", b.Obj, err)
		}
		if isDifferent(c.x, c.y, c.vx, c.vy, b) {
			t.Errorf("Ball with (%f, %f, %f, %f)) = %+v != (%f, %f, %f, %f)", c.x0, c.y0, c.vx0, c.vy0, b.Obj, c.x, c.y, c.vx, c.vy)
		}
	}
}

func TestCollide(t *testing.T) {

	// Two balls within range
	b1, _ := New(50.0, 50.0, 1.0, 1.0, 1.0, img)
	b2, _ := New(51.0, 51.0, -1.0, -1.0, 1.0, img)
	Collide(b1, b2)
	if isDifferent(50.0, 50.0, -1.0, -1.0, b1) {
		t.Errorf("Ball b1 is faulty %+v", b1.Obj)
	}
	if isDifferent(51.0, 51.0, 1.0, 1.0, b2) {
		t.Errorf("Ball b2 is faulty %+v", b2.Obj)
	}

	// Two balls out of range
	b3, _ := New(50.0, 50.0, 1.0, 1.0, 1.0, img)
	b4, _ := New(55.0, 55.0, -1.0, -1.0, 1.0, img)
	Collide(b3, b4)
	if isDifferent(50.0, 50.0, 1.0, 1.0, b3) {
		t.Errorf("Ball b3 is faulty %+v", b1.Obj)
	}
	if isDifferent(55.0, 55.0, -1.0, -1.0, b4) {
		t.Errorf("Ball b4 is faulty %+v", b2.Obj)
	}

}
