/*
Given a DNA strand, return its RNA complement (per RNA transcription).

Both DNA and RNA strands are a sequence of nucleotides.

The four nucleotides found in DNA are adenine (**A**), cytosine (**C**),
guanine (**G**) and thymine (**T**).

The four nucleotides found in RNA are adenine (**A**), cytosine (**C**),
guanine (**G**) and uracil (**U**).

Given a DNA strand, its transcribed RNA strand is formed by replacing
each nucleotide with its complement:

* `G` -> `C`
* `C` -> `G`
* `T` -> `A`
* `A` -> `U`
*/

// Package strand performs conversion from DNA to RNA
package strand

import (
	"errors"
)

// Function Complement returns the complement of a given nucleotide
func Complement(r rune) rune {
	switch r {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	default:
		errors.New("Error : The specified input is not a recognized nucleotide !")
	}
	return '0'
}

// Function toRNA converts a string sequence or DNA to its equivalent in RNA
func ToRNA(dna string) string {
	rna := []rune(dna)
	for i := 0; i < len(dna); i++ {
		rna[i] = Complement(rna[i])
	}
	return string(rna)
}
