package bounce

import (
	"github.com/Aoana/ball-sim-go/internal/pkg/ball"
	"github.com/Aoana/ball-sim-go/pkg/gfxutil"
	"github.com/Aoana/ball-sim-go/pkg/mathutil"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
)

// Simulation variables
var (
	// Mathematical values
	dt, g float64 = 10.0, 9.80665
	// Images
	backgroundImage, leftWallImage, rightWallImage *ebiten.Image
	BallImage                                      *ebiten.Image
	// Screen Resolution
	ScreenWidth  = 1600
	ScreenHeight = 900
	// Starting values for balls
	X0            = float64(ScreenWidth) / 2
	Y0            = float64(ScreenHeight) / 10
	MinV0 float64 = -50
	MaxV0 float64 = 50
)

func init() {
	// Load images
	backgroundImage, _ = gfxutil.LoadPng("./assets/sky.png")
	leftWallImage, _ = gfxutil.LoadPng("./assets/wall-left.png")
	rightWallImage, _ = gfxutil.LoadPng("./assets/wall-right.png")
	BallImage, _ = gfxutil.LoadPng("./assets/basketball.png")
}

// StartValues set starting position and velocity for a slice of balls
// Fixed starting position and velocity is random
func StartValues(nballs int) error {

	for i := 0; i < nballs; i++ {
		// Random starting velocity
		vx0, _ := mathutil.RandInRange(MinV0, MaxV0)
		vy0, _ := mathutil.RandInRange(MinV0, MaxV0)
		// Ball constructor
		err := ball.Add(X0, Y0, vx0, vy0, 0.05, BallImage)
		if err != nil {
			return err
		}
	}
	return nil
}

// DrawScenery is a helper function to draw background and walls
func DrawScenery(screen *ebiten.Image) {
	// Draw background
	gfxutil.PrintImage(screen, backgroundImage, 0, 0, 3.0, 2.3)
	// Draw walls
	gfxutil.PrintImage(screen, leftWallImage, -50, 60, 1.9, 0.9)
	gfxutil.PrintImage(screen, rightWallImage, 1400, 50, 1.7, 0.9)
}

// Timestep is a helper function to perform a timestep with position and velocity updates
func Timestep(o *objects.Object) {
	o.Position(dt)
	o.Velocity(0, g, dt)
}

// OutOfBound is a helper function to set the right boundary
// The values are simply set to fit the scenery
func OutOfBound(b *ball.Ball) {
	ball.Boundary(b, 100, float64(ScreenWidth)-60, -500, float64(ScreenHeight)-100, 0.8)
}
