// Package allyourbase performs base conversions
package allyourbase

import (
	"errors"
)

func Power(a, n int) int {
	var i, result int
	result = 1
	for i = 0; i < n; i++ {
		result *= a
	}
	return result
}

// Function ConvertToBase converts a number expressed in a input base in its expression in a new base.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	// Variables
	output := make([]int, 0)
	var inputNumber int = 0

	// Error handling : input and output base 1 case
	if inputBase <= 1 {
		output = append(output, 0)
		return output, errors.New("input base must be >= 2")
	}
	if outputBase <= 1 {
		output = append(output, 0)
		return output, errors.New("output base must be >= 2")
	}

	// Convert to base 10
	for i := 0; i < len(inputDigits); i++ {
		if inputDigits[i] < 0 || inputDigits[i] >= inputBase {
			output = append(output, 0)
			return output, errors.New("all digits must satisfy 0 <= d < input base")
		} else {
			inputNumber += inputDigits[i] * Power(inputBase, len(inputDigits)-1-i)
		}
	}

	// Input 0 case
	if inputNumber == 0 {
		output = append(output, 0)
		return output, nil
	}

	// Calculate digits in new base
	for inputNumber != 0 {
		output = append([]int{inputNumber % outputBase}, output...)
		inputNumber -= inputNumber % outputBase
		inputNumber = inputNumber / outputBase
	}
	return output, nil
}
