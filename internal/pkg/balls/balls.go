package balls

type Ball struct {
	X, Y   float64
	VX, VY float64
}

func (b *Ball) UpdatePosition(dt float64) error {
	b.X = b.X + b.VX/dt
	b.Y = b.Y + b.VY/dt
	return nil
}

func (b *Ball) UpdateVelocity(a, dt float64) error {
	b.VY = b.VY + a/dt
	return nil
}

func New(x, y, vx, vy float64) *Ball {
	b := Ball{X: x, Y: y, VX: vx, VY: vy}
	return &b
}


