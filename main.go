package main

import (
	"fmt"
)

type ball struct {
	posX, posY float64
	velX, velY float64
	radius     float64
}

func printStatus(balls ...ball) {
	fstring := "Ball %d: pos(%2.2f, %2.2f) vel(%2.2f, %2.2f)\n"
	for i, b := range balls {
		fmt.Printf(fstring, i, b.posX, b.posY, b.velX, b.velY)
	}
}

func main() {

	var nballs int

	fmt.Println("Welcome to ball simulator! Please insert an integer")
	fmt.Scanf("%d", &nballs)
	fmt.Printf("You have entered %d balls \n", nballs)
	balls := make([]ball, nballs)

	for i := range balls {
		balls[i].posX = 0
		balls[i].posY = float64(i) + 1
		balls[i].velX = 2
		balls[i].velY = 0
		balls[i].radius = 1
	}
	printStatus(balls...)

}
