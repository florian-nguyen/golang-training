// Package anagram identifies anagrams.
package anagram

import (
	"strings"
)

// Function Detect finds anagrams of a given string in a list of candidates.
func Detect(subject string, candidates []string) []string {
	output := []string{}
	var current string
	for _, c := range candidates {
		current = c

		// Candidate = subject case
		if strings.Compare(strings.ToLower(c), strings.ToLower(subject)) == 0 {
			break
		}

		// General case
		for _, letter := range []rune(strings.ToLower(subject)) {
			if strings.ContainsRune(current, letter) {
				current = strings.Replace(strings.ToLower(current), string(letter), "", 1)
			} else {
				break
			}
		}
		if strings.Compare(current, "") == 0 {
			output = append(output, c)
		}
	}
	return output
}
