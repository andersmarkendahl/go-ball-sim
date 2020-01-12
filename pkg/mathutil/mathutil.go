package mathutil

import (
	"errors"
	"math/rand"
	"time"
)

// RandInRange returns a random number in given range
func RandInRange(min, max float64) (float64, error) {

	if min >= max {
		return 0, errors.New("min less or equal to max")
	}

	rand.Seed(time.Now().UnixNano())

	return rand.Float64()*(max-min) + min, nil
}
