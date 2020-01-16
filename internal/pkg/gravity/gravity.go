package gravity

import (
	"github.com/Aoana/ball-sim-go/pkg/gfxutil"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
)

// Simulation variables
var (
	// Mathematical values
	dt, g float64 = 10.0, 9.80665
	// Images
	backgroundImage, leftWallImage, rightWallImage *ebiten.Image
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
