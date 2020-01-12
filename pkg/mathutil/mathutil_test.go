package mathutil

import (
	"testing"
)

// RandInRange returns a random number in given range
func TestRandInRange(t *testing.T) {
	// Positive tests
	for _, c := range []struct {
		min, max float64
	}{
		{0, 1},
		{-1, 0},
		{-1, 1},
		{0, 45525},
		{-0.44224, 0},
		{-0.00002, 122342.4},
	} {
		val, err := RandInRange(c.min, c.max)
		if err != nil {
			t.Errorf("RandInRange(%f, %f) -> %s", c.min, c.max, err)
		}
		if val > c.max || val < c.min {
			t.Errorf("RandInRange(%f, %f) -> %f", c.min, c.max, val)
		}
	}
	// Negative tests
	for _, c := range []struct {
		min, max float64
	}{
		{4200, 4100},
		{-3.1416, -3.1417},
	} {
		val, err := RandInRange(c.min, c.max)
		if err == nil {
			t.Errorf("RandInRange(%f, %f) -> %f", c.min, c.max, val)
		}
	}

}
