// Package wordy gives the result of an operation formulated in the form of a question in string format.
package wordy

import (
	"strconv"
	"strings"
)

// Function Answer gives the integer result of the operation stated in a question of ftype string.
func Answer(question string) (int, bool) {

	// Initial string formatting
	s := strings.Replace(question, "What is ", "", -1)
	s = strings.Replace(s, "plus", "+", -1)
	s = strings.Replace(s, "minus", "-", -1)
	s = strings.Replace(s, "multiplied by", "x", -1)
	s = strings.Replace(s, "divided by", "/", -1)
	s = strings.Replace(s, "squared", "^ 2", -1)
	s = strings.Replace(s, "cubed", "^ 3", -1)
	s = strings.TrimRight(s, "?")
	list := strings.Split(s, " ")

	// Exit if first argument is not a number or number of values incoherent
	total, ok := strconv.Atoi(string(list[0]))
	if ok != nil || len(list)%2 != 1 {
		return 0, false
	}

	// General case with operators
	for i := 1; i < len(list)/2+1; i++ {

		// Exit if value to add/multiply is not a number
		value, ok := strconv.Atoi(string(list[2*i]))
		if ok != nil {
			return 0, false
		}

		switch list[2*i-1] {
		case "+":
			total += value
		case "-":
			total -= value
		case "x":
			total *= value
		case "/":
			total /= value
		// case "^":
		// 	factor := total
		// 	for k := 0; k < value; k++ {
		// 		total *= factor
		// 	}
		default:
			return 0, false // Exit if operand not recognized
		}
	}
	return total, true
}
