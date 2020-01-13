package gravity

import (
	"github.com/Aoana/ball-sim-go/pkg/gfxutil"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
)

// Global variables
var (
	// Mathematical values
	dt, g float64 = 10.0, 9.80665
	// Screen Resolution
	ScreenWidth, ScreenHeight int = 1600, 900
	// Starting values for balls
	X0, Y0       float64 = float64(ScreenWidth) / 2, float64(ScreenHeight) / 10
	MinV0, MaxV0 float64 = -50, 50
	// Images
	backgroundImage, leftWallImage, rightWallImage *ebiten.Image
)

func init() {

	// Load images
	backgroundImage, _ = gfxutil.LoadPng("./assets/sky.png")
	leftWallImage, _ = gfxutil.LoadPng("./assets/wall-left.png")
	rightWallImage, _ = gfxutil.LoadPng("./assets/wall-right.png")

}

func DrawScenery(screen *ebiten.Image) {
	// Draw background
	gfxutil.PrintImage(screen, backgroundImage, 0, 0, 3.0, 2.3)
	// Draw walls
	gfxutil.PrintImage(screen, leftWallImage, -50, 60, 1.9, 0.9)
	gfxutil.PrintImage(screen, rightWallImage, 1400, 50, 1.7, 0.9)
}

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
