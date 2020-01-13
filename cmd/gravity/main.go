package main

import (
	"fmt"
	"github.com/Aoana/ball-sim-go/pkg/gfxutil"
	"github.com/Aoana/ball-sim-go/pkg/mathutil"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
	"log"
)

// Global variables
var (
	ballList                      []*objects.Object
	dt, g                         float64 = 10.0, 9.80665
	screenWidth, screenHeight     int     = 1600, 900
	backgroundImage               *ebiten.Image
	leftWallImage, rightWallImage *ebiten.Image
)

func timestep(o *objects.Object) {
	o.Position(dt)
	o.Velocity(0, g, dt)
}

func bounce(o *objects.Object) error {

	var factor = 0.8

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
		timestep(ballList[i])
		bounce(ballList[i])
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw background
	gfxutil.PrintImage(screen, backgroundImage, 0, 0, 3.0, 2.3)

	// Draw walls
	gfxutil.PrintImage(screen, leftWallImage, -50, 60, 1.9, 0.9)
	gfxutil.PrintImage(screen, rightWallImage, 1400, 50, 1.7, 0.9)

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

	// Load images
	ballImage, err := gfxutil.LoadPng("./assets/basketball.png")
	if err != nil {
		log.Fatal(err)
	}
	backgroundImage, err = gfxutil.LoadPng("./assets/sky.png")
	if err != nil {
		log.Fatal(err)
	}
	leftWallImage, err = gfxutil.LoadPng("./assets/wall-left.png")
	if err != nil {
		log.Fatal(err)
	}
	rightWallImage, err = gfxutil.LoadPng("./assets/wall-right.png")
	if err != nil {
		log.Fatal(err)
	}

	// Starting point for balls
	x0 := float64(screenWidth) / 2
	y0 := float64(screenHeight) / 10

	// Initialize balls
	for i := range ballList {
		// Random starting velocity
		vx0, _ := mathutil.RandInRange(-50, 50)
		vy0, _ := mathutil.RandInRange(-50, 50)
		// Object constructor
		ballList[i], err = objects.New(x0, y0, vx0, vy0, ballImage)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Run simulation loop
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Ball Sim Go"); err != nil {
		log.Fatal(err)
	}
}
