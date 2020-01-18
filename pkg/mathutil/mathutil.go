package mathutil

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

// RandInRange returns a random number in given float range
func RandInRange(min, max float64) (float64, error) {

	if min >= max {
		return 0, errors.New("min less or equal to max")
	}

	rand.Seed(time.Now().UnixNano())

	return rand.Float64()*(max-min) + min, nil
}

// Magnitude of a vector
func Magnitude(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// DotProduct (inner product) of two vectors (x1, y1) (x2, y2)
func DotProduct(x1, y1, x2, y2 float64) float64 {
	return x1*x2 + y1*y2
}
