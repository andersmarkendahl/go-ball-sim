package main

import (
	"fmt"
	"github.com/Aoana/ball-sim-go/internal/pkg/bounce"
	"github.com/Aoana/ball-sim-go/pkg/gfxutil"
	"github.com/Aoana/ball-sim-go/pkg/mathutil"
	"github.com/hajimehoshi/ebiten"
	"log"
)

// Global variables
var (
	ballList []*bounce.Ball
)

func update(screen *ebiten.Image) error {

	// Move balls, update velocity and check for bounce
	for i := range ballList {
		bounce.Timestep(ballList[i].Obj)
		bounce.Bounce(ballList[i])
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw background and walls
	bounce.DrawScenery(screen)

	// Draw balls
	for i := range ballList {
		gfxutil.PrintImage(screen, ballList[i].Image, ballList[i].Obj.X, ballList[i].Obj.Y, 0.05, 0.05)
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
	ballList = make([]*bounce.Ball, nballs)

	ballImage, err := gfxutil.LoadPng("./assets/basketball.png")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize balls
	for i := range ballList {
		// Random starting velocity
		vx0, _ := mathutil.RandInRange(bounce.MinV0, bounce.MaxV0)
		vy0, _ := mathutil.RandInRange(bounce.MinV0, bounce.MaxV0)
		// Ball constructor
		ballList[i], err = bounce.New(bounce.X0, bounce.Y0, vx0, vy0, ballImage)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Run simulation loop
	if err := ebiten.Run(update, bounce.ScreenWidth, bounce.ScreenHeight, 1, "Ball Sim Go"); err != nil {
		log.Fatal(err)
	}
}
