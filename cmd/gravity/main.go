package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	b "github.com/Aoana/ball-sim-go/internal/pkg/balls"
	"image/png"
	"log"
	"os"
)

type object interface {
	UpdatePosition(float64) error
	UpdateVelocity(float64, float64) error
}

func objectTimestep(o object) {
	o.UpdatePosition(DT)
	o.UpdateVelocity(G, DT)
}

// Global variables
var (
	ballSprite *ebiten.Image
	ballList     []*b.Ball
	DT, G      float64 = 10.0, 9.80665
)

func init() {

	// Preload images
	file, err := os.Open("./assets/basketball.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(os.Stderr, "%s: %v\n", "./assets/basketball.png", err)
	}

	ballSprite, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

}

func update(screen *ebiten.Image) error {

	// Move balls and update velocity
	for i := range ballList {
		objectTimestep(ballList[i])
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw balls
	for i := range ballList {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.05, 0.05)
		op.GeoM.Translate(ballList[i].X, ballList[i].Y)
		screen.DrawImage(ballSprite, op)
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
	ballList = make([]*b.Ball, nballs)

	// Call constructor to set initial values
	for i := range ballList {
		ballList[i] = b.New(0, float64(i)*50, 20, 0)
	}

	// Run simulation loop
	if err := ebiten.Run(update, 1600, 900, 1, "Ball Sim Go"); err != nil {
		log.Fatal(err)
	}
}
