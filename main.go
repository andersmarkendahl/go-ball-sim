package main

import (
	"fmt"
	"os"
)

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
		balls[i] = createBall(float64(i)+1, 2)
	}
	printStatus(balls...)

}
