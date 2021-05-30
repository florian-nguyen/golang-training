/*
The ISBN-10 format is 9 digits (0 to 9) plus one check character (either a digit or an X only). In the case the check character is an X, this represents the value '10'. These may be communicated with or without hyphens, and can be checked for their validity by the following formula:

```
(x1 * 10 + x2 * 9 + x3 * 8 + x4 * 7 + x5 * 6 + x6 * 5 + x7 * 4 + x8 * 3 + x9 * 2 + x10 * 1) mod 11 == 0
```

If the result is 0, then it is a valid ISBN-10, otherwise it is invalid.
*/

// Package ISBN implements the ISBN-10 verification process.
package isbn

import (
	"strings"
)

// Function IsValidISBN tells if a given code is valid or not according to the ISBN-10 verification process.
func IsValidISBN(input string) bool {
	// Removing "-" symbols
	var s string = strings.Replace(input, "-", "", -1)
	var sum int = 0
	runes := []rune(s)
	if len(s) != 10 {
		return false
	} else {
		// Check that the first 8 characters are numbers
		for i := 0; i < 10; i++ {
			if runes[i] < 48 || runes[i] > 57 {
				if i == 9 && (runes[i] == 88 || runes[i] == 120) {
					sum += 10
				} else {
					return false
				}
			} else {
				sum += int(runes[i]-'0') * (10 - i)
			}
		}
		// ISBN-10 criteria computation
		if sum%11 == 0 {
			return true
		} else {
			return false
		}
	}

}
