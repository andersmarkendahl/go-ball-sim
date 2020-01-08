package main

import (
	"fmt"
	"os"
)

var T int = 10
var DT, G float64 = 10.0 , -9.80665

type ball struct {
	X, Y   float64
	vX, vY float64
	R      float64
}

func printStatus(balls ...*ball) {
	fstring := "Ball %d: pos(%2.2f, %2.2f) vel(%2.2f, %2.2f)\n"
	for i, b := range balls {
		fmt.Printf(fstring, i, b.X, b.Y, b.vX, b.vY)
	}
}

func createBall(y, vx float64) *ball {
	b := ball{Y: y, vX: vx}
	b.R = 2
	return &b
}

func (b *ball) updatePosition() error {
	b.X = b.X + b.vX/DT
	b.Y = b.Y + b.vY/DT
	return nil
}

func (b *ball) updateVelocity() error {
	b.vY = b.vY + G/DT
	return nil
}

type object interface {
	updatePosition() error
	updateVelocity() error
}

func objectTimestep (o object) {
	o.updatePosition()
	o.updateVelocity()
}

func main() {

	var nballs int

	fmt.Println("Welcome to ball simulator! Please insert an integer")
	_, err := fmt.Scanf("%d", &nballs)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	fmt.Printf("You have entered %d balls \n", nballs)
	balls := make([]*ball, nballs)

	for i := range balls {
		balls[i] = createBall(float64(i) * 5, 2)
	}
	printStatus(balls...)

    for t := 0; t <= T; t++ {
		for i := range balls {
			objectTimestep(balls[i])
		}
		fmt.Println(">>> TIME: ", t)
		printStatus(balls...)
    }

}
