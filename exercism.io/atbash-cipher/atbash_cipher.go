// Package Atbash performs a reverse cipher algorithm.
package atbash

import ()

// Function Atbash returns the result of the Atbash cipher.
func Atbash(input string) string {
	r := []rune{}
	for _, value := range input {
		var character rune
		if value >= '0' && value <= '9' {
			character = value
		} else if value >= 'A' && value <= 'Z' {
			character = 'z' - value + 'A'
		} else if value >= 'a' && value <= 'z' {
			character = 'z' - value + 'a'
		} else {
			// Important to deal with other characters
			continue
		}

		if len(r)%6 == 5 {
			r = append(r, ' ')
		}
		r = append(r, character)
	}
	return string(r)

}
