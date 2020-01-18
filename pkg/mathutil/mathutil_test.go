package mathutil

import (
	"testing"
)

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

func TestMagnitude(t *testing.T) {

	// Positive tests
	for _, c := range []struct {
		x, y, d float64
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{3.0, 4.0, 5},
		{-3.0, 4.0, 5},
		{3.0, -4.0, 5},
		{-3.0, -4.0, 5},
	} {
		got := Magnitude(c.x, c.y)
		if got != c.d {
			t.Errorf("Magnitude(%f, %f) got %f want %f", c.x, c.y, got, c.d)
		}
	}
}

func TestDotProduct(t *testing.T) {

	// Positive tests
	for _, c := range []struct {
		x1, y1, x2, y2, d float64
	}{
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1},
		{0, 1, 0, 1, 1},
		{2.0, 2.5, 1.5, 2.0, 8.0},
		{34.4, 34.4, 43.3, -43.3, 0},
	} {
		got := DotProduct(c.x1, c.y1, c.x2, c.y2)
		if got != c.d {
			t.Errorf("Magnitude(%f, %f, %f, %f) got %f want %f", c.x1, c.y1, c.x2, c.y2, got, c.d)
		}
	}
}
