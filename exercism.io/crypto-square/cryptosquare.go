// Package cryptoquare generates messages encoded by means of the square code.
package cryptosquare

/*
Implement the classic method for composing secret messages called a square code.

Given an English text, output the encoded version of that text.

First, the input is normalized: the spaces and punctuation are removed
from the English text and the message is downcased.

Then, the normalized characters are broken into rows.  These rows can be
regarded as forming a rectangle when printed with intervening newlines.
*/

import (
	"math"
	"strings"
)

func Encode(input string) (output string) {
	// First step : normalization
	s := strings.ToLower(input)
	for _, r := range s {
		switch {
		case r >= 32 && r <= 47, r >= 58 && r <= 64, r >= 91 && r <= 96, r >= 123 && r <= 126:
			s = strings.Replace(s, string(r), "", -1)
		}
	}

	// Second step : computation of c and r
	s, c, r := Rectangle(s)

	// Third step : Compute transpose matrix
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			output += string(s[j*c+i])
			if j == r-1 && i != c-1 {
				output += " "
			}
		}
	}
	return output

}

// Function Rectangle returns r and c, and a string of length r x c (space padding)
func Rectangle(s string) (output string, c, r int) {
	// Empty string case
	if len(s) == 0 {
		return s, c, r
	}

	// r and c computation
	c = int(math.Ceil(math.Sqrt(float64(len(s)))))
	r = int(math.Floor(math.Sqrt(float64(len(s)))))
	if r*c < len(s) {
		r++
	}

	// Adding space padding
	output = s
	for i := 0; i < c*r-len(s); i++ {
		output += " "
	}
	return output, c, r
}
