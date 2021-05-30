// Package phonenumber formats phone numbers according to the NANP standard.
package phonenumber

import (
	"errors"
)

// Function Number returns the formated number
func Number(input string) (string, error) {
	// Removing non-numeral characters
	numbers := make([]rune, 0)
	for _, r := range []rune(input) {
		if r >= 48 && r <= 57 {
			numbers = append(numbers, r)
		}
	}

	// Number length inferior to 10
	if len(numbers) <= 9 {
		return "", errors.New("Error: 9 digit input.")
	}

	// Number length superior to 11
	if len(numbers) > 11 {
		return "", errors.New("Error: more than 11 digit input.")
	}

	// Number length of 11
	if len(numbers) == 11 && numbers[0] == '1' {
		numbers = numbers[1:]
	} else if len(numbers) == 11 && numbers[0] != '1' {
		return "", errors.New("Error: 11 digit input with first number not 1.")
	}

	// Finally, confirm validity of number format
	if numbers[0]-'0' >= 2 && numbers[3]-'0' >= 2 {
		return string(numbers), nil
	} else {
		return "", errors.New("Error: First and fourth digits should be higher than 2.")
	}

}

// Function Format returns a formatted version of the input phone number.
func Format(s string) (string, error) {
	number, err := Number(s)
	if err != nil {
		return "", err
	} else {
		runes := []rune(number)
		return "(" + string(runes[0:3]) + ") " + string(runes[3:6]) + "-" + string(runes[6:10]), nil
	}
}

// Function AreaCode returns the area code associated to a given phone number.
func AreaCode(s string) (string, error) {
	number, err := Number(s)
	if err != nil {
		return "", err
	}
	r := []rune(number)
	return string(r[0:3]), nil
}
