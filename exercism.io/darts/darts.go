// Package darts calculates darts scores.
package darts

import (
	"math"
)

const (
	INNER_RADIUS  float64 = 1
	MIDDLE_RADIUS float64 = 5
	OUTER_RADIUS  float64 = 10
)

// Function Score returns the score associated to a (x,y) position.
func Score(x, y float64) int {
	radius := math.Sqrt(x*x + y*y)
	switch a := radius; {
	case a <= INNER_RADIUS:
		return 10
	case a <= MIDDLE_RADIUS:
		return 5
	case a <= OUTER_RADIUS:
		return 1
	default:
		return 0
	}
}
