package main

import (
	"flag"
	"github.com/Aoana/ball-sim-go/internal/pkg/ball"
	"github.com/Aoana/ball-sim-go/internal/pkg/collision"
	"github.com/hajimehoshi/ebiten"
	"log"
)

func update(screen *ebiten.Image) error {

	// Logical update
	for i := range ball.BallList {
		// Check balls against collision, ignore self and already checked
		for j := range ball.BallList {
			if i > j {
				ball.Collide(ball.BallList[i], ball.BallList[j])
			}
		}
		// Move balls
		collision.Timestep(ball.BallList[i])
		// Check if reached goal
		collision.Goal(ball.BallList[i])
		// Check bounce on walls
		collision.Edge(ball.BallList[i])
	}

	// Remove deactivated balls
	l := len(ball.BallList)
	for i := 0; i < l; i++ {
		if !ball.BallList[i].Active {
			err := ball.Remove(i)
			if err != nil {
				return err
			}
			l--
		}
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw background and walls
	collision.DrawScenery(screen)
	// Draw balls
	for i := range ball.BallList {
		ball.Print(screen, ball.BallList[i])
	}
	return nil
}

func main() {

	var err error

	// Check user specified number of balls
	nballs := flag.Int("nballs", 10, "Number of balls")
	flag.Parse()

	// Initialize balls
	err = collision.StartValues(*nballs)
	if err != nil {
		log.Fatal(err)
	}

	// Run simulation loop
	if err := ebiten.Run(update, collision.ScreenWidth, collision.ScreenHeight, 1, "Ball Sim Go"); err != nil {
		log.Fatal(err)
	}
}
