// Package hamming calculates the Hamming difference between two DNA strands.
package hamming

// Import packages.
import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// Distance calculates the Hamming difference between two DNA strands specified as strings.
func Distance(a, b string) (int, error) {

	// Two input strings having different lengths results in an error and value -1.
	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return -1, errors.New(fmt.Sprintf("Error: The two sequences do not have the same length!"))
	}

	// General case: indexes are compared individually.
	var count int
	for index := 0; index < utf8.RuneCountInString(a); index++ {
		if a[index] != b[index] {
			count++
		}
	}
	return count, nil

}
