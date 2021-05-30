// Package isogram finds out if a word or a sentence is an isogram, i.e. a word or phrase without a repeating letter.
package isogram

// Import packages.
import (
	"strings"
	"unicode/utf8"
)

// IsIsogram returns TRUE if the specified input is an isogram and FALSE otherwise.
func IsIsogram(s string) bool {

	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)

	for i := 0; i < utf8.RuneCountInString(s); i++ {
		for j := i + 1; j < utf8.RuneCountInString(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true

}
