package gravity

import (
	"github.com/Aoana/ball-sim-go/pkg/objects"
)

// Global variables
var (
	dt, g                     float64 = 10.0, 9.80665
	ScreenWidth, ScreenHeight int     = 1600, 900
)

func Timestep(o *objects.Object) {
	o.Position(dt)
	o.Velocity(0, g, dt)
}

func Bounce(o *objects.Object) error {

	var factor = 0.8

	if o.Y > float64(ScreenHeight-100) {
		if o.VY > 0 {
			o.VY = -o.VY * factor
		}
	}

	if o.X > float64(ScreenWidth-100) {
		if o.VX > 0 {
			o.VX = -o.VX * factor
		}
	}

	if o.X < 100 {
		if o.VX < 0 {
			o.VX = -o.VX * factor
		}
	}
	return nil
}
