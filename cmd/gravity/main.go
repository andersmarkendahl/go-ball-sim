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
	ballList                  []*objects.Object
	dt, g                     float64 = 10.0, 9.80665
	screenWidth, screenHeight int     = 1600, 900
	backgroundImage           *ebiten.Image
)

type timestep interface {
	Position(float64) error
	Velocity(float64, float64, float64) error
}

func performTimestep(t timestep) {
	t.Position(dt)
	t.Velocity(0, g, dt)
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
		performTimestep(ballList[i])
		bounce(ballList[i])
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw background
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.0, 0.8)
	op.GeoM.Translate(0, 0)
	screen.DrawImage(backgroundImage, op)

	// Draw balls
	for i := range ballList {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.02, 0.02)
		op.GeoM.Translate(ballList[i].X, ballList[i].Y)
		screen.DrawImage(ballList[i].Image, op)
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
	fmt.Printf("Start allocating %d balls...\n", nballs)
	// Create a slice of number of balls
	ballList = make([]*objects.Object, nballs)
	fmt.Printf("Done allocating %d balls...\n", nballs)

	// Load images
	ballImage, err := gfxutil.LoadPng("./assets/basketball.png")
	if err != nil {
		log.Fatal(err)
	}
	backgroundImage, err = gfxutil.LoadPng("./assets/sky.png")
	if err != nil {
		log.Fatal(err)
	}

	x0 := float64(screenWidth) / 2
	y0 := float64(screenHeight) / 10

	// Call constructor to set initial values
	fmt.Println("Start setting values of balls")
	// Create a slice of number of balls
	for i := range ballList {
		vx0, errX := mathutil.RandInRange(-50, 50)
		if errX != nil {
			log.Fatal(errX)
		}
		vy0, errY := mathutil.RandInRange(-50, 50)
		if errY != nil {
			log.Fatal(errY)
		}
		ballList[i], err = objects.New(x0, y0, vx0, vy0, ballImage)
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
