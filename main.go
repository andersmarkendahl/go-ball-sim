package main

import (
	"fmt"
)

type ball struct {
	posX, posY float64
	velX, velY float64
	radius     float64
}

func main() {

	var nballs int

	fmt.Println("Welcome to ball simulator! Please insert an integer")

	fmt.Scanf("%d", &nballs)
	fmt.Printf("You have entered %d balls \n", nballs)

	balls := make([]ball, nballs)
	fmt.Println("Empty Balls:", balls)

	for i := range balls {
		balls[i].posX = 0
		balls[i].posY = float64(i) + 1
		balls[i].velX = 2
		balls[i].velY = 0
		balls[i].radius = 1
	}
	fmt.Println("Set Balls:", balls)

}
