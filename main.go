package main

import (
	"fmt"
	"os"
	"log"
	"image/png"
	"github.com/hajimehoshi/ebiten"
)


type ball struct {
	X, Y   float64
	vX, vY float64
}

func printStatus(balls ...*ball) {
	fstring := "Ball %d: pos(%2.2f, %2.2f) vel(%2.2f, %2.2f)\n"
	for i, b := range balls {
		fmt.Printf(fstring, i, b.X, b.Y, b.vX, b.vY)
	}
}

func createBall(y, vx float64) *ball {
	b := ball{Y: y, vX: vx}
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

var (
	ballSprite	*ebiten.Image
	balls		[]*ball
	DT, G float64 = 10.0 , 9.80665
)

func init() {

	// Preload images
	file, err := os.Open("./images/basketball.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(os.Stderr, "%s: %v\n", "./images/basketball.png", err)
	}

    ballSprite, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

}

func update(screen *ebiten.Image) error {

	for i := range balls {
		objectTimestep(balls[i])
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	for i := range balls {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.05, 0.05)
		op.GeoM.Translate(balls[i].X, balls[i].Y)
	    screen.DrawImage(ballSprite, op)
	}
	return nil
}
func main() {

	var nballs int

	fmt.Println("Welcome to ball simulator! Please insert an integer")
	_, err := fmt.Scanf("%d", &nballs)
	if err != nil {
		log.Fatal(err)
	}
	balls = make([]*ball, nballs)

	for i := range balls {
		balls[i] = createBall(float64(i) * 50, 20)
	}

	if err := ebiten.Run(update, 1600, 900, 1, "Ball Sim Go"); err != nil {
		log.Fatal(err)
	}
}
