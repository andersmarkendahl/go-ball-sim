package collision

import (
	"github.com/Aoana/ball-sim-go/internal/pkg/ball"
	"github.com/Aoana/ball-sim-go/pkg/gfxutil"
	"github.com/Aoana/ball-sim-go/pkg/mathutil"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
	"math"
)

// Simulation variables
var (
	// Mathematical values
	dt = 10.0
	// Images
	backgroundImage *ebiten.Image
	SoccerBallImage *ebiten.Image
	// Screen Resolution
	ScreenWidth  = 1600
	ScreenHeight = 900
	// Starting values for balls
	//	X0            = float64(ScreenWidth) / 2
	//	Y0            = float64(ScreenHeight) / 10
	MinV0 float64 = -50
	MaxV0 float64 = 50
)

func init() {

	// Load background image
	backgroundImage, _ = gfxutil.LoadPng("./assets/soccerfield.png")
	SoccerBallImage, _ = gfxutil.LoadPng("./assets/soccerball.png")

}

// StartValues set starting position and velocity for a slice of balls
// Positions spread in a square and velocity is random
func StartValues(bs []*ball.Ball) error {

	var err error
	l := len(bs)
	s := int(math.Round(1 + math.Sqrt(float64(l))))
	cx := make([]float64, 0)
	cy := make([]float64, 0)

	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			cx = append(cx, float64(300+i*20))
			cy = append(cy, float64(300+j*20))
		}
	}

	for i := range bs {
		// Random starting velocity
		vx0, _ := mathutil.RandInRange(MinV0, MaxV0)
		vy0, _ := mathutil.RandInRange(MinV0, MaxV0)
		// Ball constructor
		bs[i], err = ball.New(cx[i], cy[i], vx0, vy0, 0.07, SoccerBallImage)
		if err != nil {
			return err
		}
	}
	return nil
}

// DrawScenery is a helper function to draw background
func DrawScenery(screen *ebiten.Image) {
	// Draw background
	gfxutil.PrintImage(screen, backgroundImage, 0, 0, 4.0, 3.5)
}

// Timestep is a helper function to perform a timestep with position and velocity updates
func Timestep(o *objects.Object) {
	o.Position(dt)
}

// OutOfBound is a helper function to set the right boundary
func OutOfBound(b *ball.Ball) {
	ball.Boundary(b, 0, float64(ScreenWidth), 0, float64(ScreenHeight), float64(1))
}
