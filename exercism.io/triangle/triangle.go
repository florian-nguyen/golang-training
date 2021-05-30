/*
Determine if a triangle is equilateral, isosceles, or scalene.

An equilateral triangle has all three sides the same length.

An isosceles triangle has at least two sides the same length. (It is sometimes
specified as having exactly two sides the same length, but for the purposes of
this exercise we'll say at least two.)

A scalene triangle has all sides of different lengths.
*/

// Package Triangle determines the nature of a triangle.
package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

// Function IsTriangle determines if three lengths a,b,c can constitute a triangle
func IsTriangle(a, b, c float64) bool {
	min := math.Min(math.Min(a, b), c)
	max := math.Max(math.Max(a, b), c)
	mid := a + b + c - min - max
	if max <= min+mid && min > 0 {
		return true
	} else {
		return false
	}
}

// Function KindFromSides determines the kind of a triangle based on its edges' length.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if !IsTriangle(a, b, c) {
		k = NaT
	} else if a == c && a == b {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else if a != b && a != c && b != c {
		k = Sca
	}
	return k
}
