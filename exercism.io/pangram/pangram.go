// Package pangram identifies Pangrams.
package pangram

import (
	"strings"
)

// Function IsPangram determines whether a given string is a Pangram or not.
func IsPangram(s string) bool {
	for i := 97; i < 123; i++ {
		if !strings.ContainsRune(strings.ToLower(s), rune(i)) {
			return false
		}
	}
	return true
}
