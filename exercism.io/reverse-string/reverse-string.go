// Package reverse is employed to reverse strings.
package reverse

// Reverse computes a reversed version of the inputed string.
func Reverse(s string) string {

	runes := []rune(s) // Array of runes composing the input string s.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Reverse the rune order.
	}
	return string(runes) // Recompose as a string.
}
