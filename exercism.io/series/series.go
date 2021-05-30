/*
Given a string of digits, output all the contiguous substrings of length `n` in
that string in the order that they appear.

For example, the string "49142" has the following 3-digit series:

- "491"
- "914"
- "142"

And the following 4-digit series:

- "4914"
- "9142"

And if you ask for a 6-digit series from a 5-digit string, you deserve
whatever you get.
*/

// Package Series allows the computation of series of consecutive substrings
package series

import ()

// Function All returns a list of all n-long substrings.
func All(n int, input string) []string {
	var output []string
	for i := 0; i < len(input)-n+1; i++ {
		output = append(output, input[i:i+n])
	}
	return output
}

// Function UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, input string) string {
	var max int
	if n > len(input) {
		max = len(input)
	} else {
		max = n
	}
	return input[0:max]
}

// Function First does the same as UnsafeFirst but takes into account possible error cases.
func First(n int, input string) (string, bool) {
	if n > len(input) {
		return UnsafeFirst(n, input), false
	} else {
		return UnsafeFirst(n, input), true
	}
}
