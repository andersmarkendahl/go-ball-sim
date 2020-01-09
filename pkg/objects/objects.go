package objects

import (
    "github.com/hajimehoshi/ebiten"
    "image/png"
    "os"
)

type Object struct {
    X, Y   float64
    VX, VY float64
    Image  *ebiten.Image
}

func (o *Object) UpdatePosition(dt float64) error {
    o.X = o.X + o.VX/dt
    o.Y = o.Y + o.VY/dt
    return nil
}

func (o *Object) UpdateVelocity(a, dt float64) error {
    o.VY = o.VY + a/dt
    return nil
}

func New(x, y, vx, vy float64, pngimage string) (*Object, error) {

    o := Object{X: x, Y: y, VX: vx, VY: vy}

    // Load image
    file, err := os.Open(pngimage)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    img, err := png.Decode(file)
    if err != nil {
        return nil, err
    }

    o.Image, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

    return &o, nil
}

