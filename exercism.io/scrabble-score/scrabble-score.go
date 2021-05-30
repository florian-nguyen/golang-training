// Package scrabble computes the scrabble score for a word that is specified.
package scrabble

// Import packages
import (
	"strings"
	"unicode/utf8"
)

// Score calculates the Scrabble score associated to the input string.
func Score(s string) int {
	var score int
	s = strings.ToUpper(s)

	for i := 0; i < utf8.RuneCountInString(s); i++ {
		switch s[i] {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score += 1
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		default:
			score += 0
		}
	}
	return score
}
