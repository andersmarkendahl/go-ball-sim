package main

import (
	"fmt"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
	"log"
)

// Global variables
var (
	ballList                  []*objects.Object
	DT, G                     float64 = 10.0, 9.80665
	screenWidth, screenHeight int     = 1600, 900
)

type timestep interface {
	Position(float64) error
	Velocity(float64, float64) error
}

func performTimestep(t timestep) {
	t.Position(DT)
	t.Velocity(G, DT)
}

func Bounce(o *objects.Object) error {

	var factor float64 = 0.8

	if o.Y > float64(screenHeight-100) {
		if o.VY > 0 {
			o.VY = -o.VY * factor
		}
	}

	if o.X > float64(screenWidth-100) {
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

func update(screen *ebiten.Image) error {

	// Move balls and update velocity
	for i := range ballList {
		performTimestep(ballList[i])
		Bounce(ballList[i])
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw balls
	for i := range ballList {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.05, 0.05)
		op.GeoM.Translate(ballList[i].X, ballList[i].Y)
		screen.DrawImage(ballList[i].Image, op)
	}
	return nil
}

func makeRange(n int, max float64) []float64 {
	s := max / float64(n)
	r := make([]float64, n)
	for i := range r {
		r[i] = float64(i) * s
	}
	return r
}

func main() {

	// User insert number of balls
	var nballs int
	fmt.Println("Welcome to ball simulator! Please insert an integer")
	_, err := fmt.Scanf("%d", &nballs)
	if err != nil {
		log.Fatal(err)
	}
	// Create a slice of number of balls
	ballList = make([]*objects.Object, nballs)

	// Create a slice of start velocity, maximum 100
	ballVelocity := makeRange(nballs, 100)

	// Call constructor to set initial values
	for i := range ballList {
		ballList[i], err = objects.New(100, 100, ballVelocity[i], -ballVelocity[i], "./assets/basketball.png")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Run simulation loop
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Ball Sim Go"); err != nil {
		log.Fatal(err)
	}
}
