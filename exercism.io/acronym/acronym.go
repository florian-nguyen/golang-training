// Package acronym converts a phrase to its acronym.
package acronym

// Import packages.
import (
	"strings"
)

// Abbreviate generates the acronym of any string input.
func Abbreviate(s string) string {

	// Remove trailing and leading spaces, and then replace all hyphens by spaces in the input string.
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "-", " ", -1)

	// Extract and concatenate the first character of each word before converting to uppercase.
	var acronym string
	for _, word := range strings.Split(s, " ") {
		acronym = acronym + string(word[0])
	}
	acronym = strings.ToUpper(acronym)

	return acronym
}
