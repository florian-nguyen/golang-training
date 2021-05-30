// Package lsproduct stands for Least Series Product.
package lsproduct

import (
	"errors"
)

// Function calculates the largest product for a contiguous substring of digits of length n, given an input string of digits.
func LargestSeriesProduct(s string, span int) (int64, error) {

	// Check input string
	for _, rune := range []rune(s) {
		if rune < '0' || rune > '9' {
			return -1, errors.New("Error : Input string must contain only numbers !")
		}
	}

	// Check span parameter
	if span < 0 {
		return -1, errors.New("Error : Span parameter must be a positive integer !")
	}

	// Check coherence between span parameter and string length
	if len(s) < span {
		return -1, errors.New("Error : Span parameter must be inferior to input string length !")
	}

	// Else...
	var max_value int64 = 0
	var n int = len(s)
	runes := []rune(s)
	var product_i int64 = 1
	for i := 0; i < n-span+1; i++ {
		product_i = 1
		for j := 0; j < span; j++ {
			product_i = product_i * int64(runes[i+j]-'0')
		}
		if product_i > max_value {
			max_value = product_i
		}
	}

	return max_value, nil
}
