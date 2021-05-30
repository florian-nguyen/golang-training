package romannumerals

import (
	"errors"
	"strconv"
)

var ROMAN_BASE = map[int]string{0: "I", 1: "V", 2: "X", 3: "L", 4: "C", 5: "D", 6: "M"}

// ToRomanNumeral returns the input number's Roman expression
func ToRomanNumeral(n int) (string, error) {

	// Error cases
	if n <= 0 || n > 3000 {
		return "", errors.New("Error: Input number must be between 0 and 3000!")
	}

	var current int
	var s, first, second, third string
	runes := []rune(strconv.Itoa(n))

	for i := 0; i < len(runes); i++ {
		current = int(runes[i] - '0')
		first = ROMAN_BASE[2*(len(runes)-1-i)]
		second = ROMAN_BASE[2*(len(runes)-1-i)+1]
		third = ROMAN_BASE[2*(len(runes)-i)]
		switch current {
		case 1:
			s += first
		case 2:
			s += first + first
		case 3:
			s += first + first + first
		case 4:
			s += first + second
		case 5:
			s += second
		case 6:
			s += second + first
		case 7:
			s += second + first + first
		case 8:
			s += second + first + first + first
		case 9:
			s += first + third
		default:
			// if zero do nothing
		}
	}
	return s, nil
}
