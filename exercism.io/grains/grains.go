// Package grains calculates how many grains are in a given square of a chess board and gives the total number of grains on the chess board.
package grains

// Import packages.
import (
	"errors"
	"fmt"
	"math"
)

// Square calculates the number of grains in the nth square.
func Square(n int) (uint64, error) {

	if n <= 0 || n > 64 {
		return uint64(0), errors.New(fmt.Sprint("Error in the specified input number."))
	}
	return uint64(math.Pow(2, float64(n-1))), nil

}

// Total is the number of grains on the chess board.
func Total() uint64 {

	var sum uint64
	for i := 1; i <= 64; i++ {
		sum += uint64(math.Pow(2, float64(i-1)))
	}
	return sum

}
