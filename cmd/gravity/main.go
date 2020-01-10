package main

import (
	"fmt"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
	"log"
	"math/rand"
	"time"
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
		op.GeoM.Scale(0.02, 0.02)
		op.GeoM.Translate(ballList[i].X, ballList[i].Y)
		screen.DrawImage(ballList[i].Image, op)
	}
	return nil
}

func randInRange(min, max float64) float64 {

	rand.Seed(time.Now().UnixNano())

	return rand.Float64()*(max - min) + min
}


func main() {

	// User insert number of balls
	var nballs int
	fmt.Println("Welcome to ball simulator! Please insert an integer")
	_, err := fmt.Scanf("%d", &nballs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Start allocating %d balls...\n", nballs)
	// Create a slice of number of balls
	ballList = make([]*objects.Object, nballs)
	fmt.Printf("Done allocating %d balls...\n", nballs)

	// Call constructor to set initial values
	fmt.Println("Start setting values of balls")
	// Create a slice of number of balls
	for i := range ballList {
		ballList[i], err = objects.New(float64(screenWidth)/2, float64(screenHeight)/10, randInRange(-50, 50), randInRange(-50, 50), "./assets/basketball.png")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Done setting values of balls")

	// Run simulation loop
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Ball Sim Go"); err != nil {
		log.Fatal(err)
	}
}
