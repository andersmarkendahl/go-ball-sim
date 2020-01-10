package mathutil

import (
	"math/rand"
	"time"
)

// RandInRange returns a random number in given range
func RandInRange(min, max float64) float64 {

	rand.Seed(time.Now().UnixNano())

	return rand.Float64()*(max-min) + min
}
