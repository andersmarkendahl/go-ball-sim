package main

import (
	"fmt"
	"github.com/Aoana/ball-sim-go/internal/pkg/gravity"
	"github.com/Aoana/ball-sim-go/pkg/gfxutil"
	"github.com/Aoana/ball-sim-go/pkg/mathutil"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
	"log"
)

// Global variables
var (
	ballList []*objects.Object
)

func update(screen *ebiten.Image) error {

	// Move balls, update velocity and check for bounce
	for i := range ballList {
		gravity.Timestep(ballList[i])
		gravity.Bounce(ballList[i])
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw background and walls
	gravity.DrawScenery(screen)

	// Draw balls
	for i := range ballList {
		gfxutil.PrintImage(screen, ballList[i].Image, ballList[i].X, ballList[i].Y, 0.05, 0.05)
	}
	return nil
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

	ballImage, err := gfxutil.LoadPng("./assets/basketball.png")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize balls
	for i := range ballList {
		// Random starting velocity
		vx0, _ := mathutil.RandInRange(gravity.MinV0, gravity.MaxV0)
		vy0, _ := mathutil.RandInRange(gravity.MinV0, gravity.MaxV0)
		// Object constructor
		ballList[i], err = objects.New(gravity.X0, gravity.Y0, vx0, vy0, ballImage)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Run simulation loop
	if err := ebiten.Run(update, gravity.ScreenWidth, gravity.ScreenHeight, 1, "Ball Sim Go"); err != nil {
		log.Fatal(err)
	}
}
