package main

import (
	"fmt"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
	"log"
)

// Global variables
var (
	ballList []*objects.Object
	DT, G    float64 = 10.0, 9.80665
)

type timestep interface {
	UpdatePosition(float64) error
	UpdateVelocity(float64, float64) error
}

func performTimestep(t timestep) {
	t.UpdatePosition(DT)
	t.UpdateVelocity(G, DT)
}

func update(screen *ebiten.Image) error {

	// Move balls and update velocity
	for i := range ballList {
		performTimestep(ballList[i])
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

func main() {

	var nballs int

	// User insert number of balls
	fmt.Println("Welcome to ball simulator! Please insert an integer")
	_, err := fmt.Scanf("%d", &nballs)
	if err != nil {
		log.Fatal(err)
	}
	// Create a slice of number of balls
	ballList = make([]*objects.Object, nballs)

	// Call constructor to set initial values
	for i := range ballList {
		ballList[i], err = objects.New(0, float64(i)*50, 20, 0, "./assets/basketball.png")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Run simulation loop
	if err := ebiten.Run(update, 1600, 900, 1, "Ball Sim Go"); err != nil {
		log.Fatal(err)
	}
}
