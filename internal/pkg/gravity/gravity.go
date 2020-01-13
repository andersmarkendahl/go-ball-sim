package gravity

import (
	"github.com/Aoana/ball-sim-go/pkg/objects"
)

// Global variables
var (
	dt, g float64 = 10.0, 9.80665
)

func Timestep(o *objects.Object) {
	o.Position(dt)
	o.Velocity(0, g, dt)
}
