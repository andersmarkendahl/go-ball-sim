package ball

import (
	"github.com/Aoana/ball-sim-go/assets/images"
	"github.com/Aoana/ball-sim-go/pkg/gfxutil"
	"testing"
)

func isPresent(x, y, vx, vy, scale float64) bool {
	for i := range BallList {
		b := BallList[i]
		if x == b.Obj.X[0] && y == b.Obj.X[1] && vx == b.Obj.V[0] && vy == b.Obj.V[1] && scale == b.Scale {
			return true
		}
	}
	return false
}

func TestList(t *testing.T) {

	img, _ := gfxutil.LoadPngSlice(images.ImageBasketBall)

	for _, c := range []struct {
		x, y, vx, vy, scale float64
	}{
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
