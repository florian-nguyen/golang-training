// Package raindrops converts a number to a string, the contents of which depend on the number's factors.
package raindrops

// Import packages.
import (
	"strconv"
)

// Convert inputs int number to string depending on its factors being 3, 5, 7, several or none of the latter.
func Convert(n int) string {
	var s string

	if n%3 != 0 && n%5 != 0 && n%7 != 0 {
		return strconv.Itoa(n)
	}

	if n%3 == 0 {
		s += "Pling"
	}

	if n%5 == 0 {
		s += "Plang"
	}

	if n%7 == 0 {
		s += "Plong"
	}

	return s

}
