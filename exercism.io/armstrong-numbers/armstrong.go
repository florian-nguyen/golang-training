// Package armstrong determines if a given number is an Armstrong number or not.
package armstrong

import (
	"log"
	"strconv"
)

// Function Power elevates integers to their nth power.
func Power(a, n int) int {
	var i, result int
	result = 1
	for i = 0; i < n; i++ {
		result *= a
	}
	return result
}

// Function IsNumber returns true if the input integer is an Armstrong number, and false otherwise.
func IsNumber(input int) bool {
	runes := []rune(strconv.Itoa(input))
	var n int = len(runes)
	var sum int = 0

	for i := 0; i < n; i++ {
		r, err := strconv.Atoi(string(runes[i]))
		if err != nil {
			log.Fatalln(err)
		}
		sum += Power(r, n)
	}
	if sum == input {
		return true
	}
	return false
}
