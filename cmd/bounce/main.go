package main

import (
	"flag"
	"github.com/Aoana/ball-sim-go/internal/pkg/ball"
	"github.com/Aoana/ball-sim-go/internal/pkg/bounce"
	"github.com/hajimehoshi/ebiten"
	"log"
)

func update(screen *ebiten.Image) error {

	// Move balls, update velocity and check for bounce
	for i := range ball.BallList {
		bounce.Timestep(ball.BallList[i])
		bounce.OutOfBound(ball.BallList[i])
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw background and walls
	bounce.DrawScenery(screen)

	// Draw balls
	for i := range ball.BallList {
		ball.Print(screen, ball.BallList[i])
	}
	return nil
}

func main() {

	var err error

	// Check user specified number of balls
	nballs := flag.Int("nballs", 3, "Number of balls")
	flag.Parse()

	// Initialize balls
	err = bounce.StartValues(*nballs)
	if err != nil {
		log.Fatal(err)
	}

	// Run simulation loop
	if err = ebiten.Run(update, bounce.ScreenWidth, bounce.ScreenHeight, 1, "Ball Sim Go"); err != nil {
		log.Fatal(err)
	}
}
