package main

import (
	"flag"
	"github.com/Aoana/ball-sim-go/internal/pkg/ball"
	"github.com/Aoana/ball-sim-go/internal/pkg/collision"
	"github.com/Aoana/ball-sim-go/pkg/mathutil"
	"github.com/hajimehoshi/ebiten"
	"log"
)

// Global variables
var (
	ballList []*ball.Ball
)

func update(screen *ebiten.Image) error {

	// Logical update
	for i := range ballList {
		// Move balls, update velocity and check for bounce
		for j := range ballList {
			if i > j {
				ball.Collide(ballList[i], ballList[j])
			}
		}
		// Move balls
		collision.Timestep(ballList[i].Obj)
		// Check bounce on walls
		collision.OutOfBound(ballList[i])
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw background and walls
	collision.DrawScenery(screen)

	// Draw balls
	for i := range ballList {
		ball.Print(screen, ballList[i])
	}
	return nil
}

func main() {

	var err error

	// Check user specified number of balls
	nballs := flag.Int("nballs", 3, "Number of balls")
	flag.Parse()

	// Create a slice of number of balls
	ballList = make([]*ball.Ball, *nballs)

	// Initialize balls
	for i := range ballList {
		// Random starting velocity
		vx0, _ := mathutil.RandInRange(collision.MinV0, collision.MaxV0)
		vy0, _ := mathutil.RandInRange(collision.MinV0, collision.MaxV0)
		// Ball constructor
		ballList[i], err = ball.New(float64(i*5), float64(i*100), vx0, vy0, 0.1, collision.SoccerBallImage)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Run simulation loop
	if err := ebiten.Run(update, collision.ScreenWidth, collision.ScreenHeight, 1, "Ball Sim Go"); err != nil {
		log.Fatal(err)
	}
}
