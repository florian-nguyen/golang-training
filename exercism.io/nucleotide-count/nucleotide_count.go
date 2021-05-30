package dna

import (
	"errors"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (h Histogram, err error) {
	h = Histogram{'G': 0, 'A': 0, 'T': 0, 'C': 0}
	for _, char := range d {
		switch char {
		case 'G', 'A', 'T', 'C':
			h[char]++
			break
		default:
			return h, errors.New("Error: Input format incorrect!")
		}
	}

	return h, err
}
