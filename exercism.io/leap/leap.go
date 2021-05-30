/*
A leap year in the Gregorian calendar occurs:
- on every year that is evenly divisible by 4
- except every year that is evenly divisible by 100
- unless the year is also evenly divisible by 400

For example, 1997 is not a leap year, but 1996 is.  1900 is not a leap
year, but 2000 is.
*/

// Package leap tells if a given year is a leap year or not.
package leap

// Function IsLeapYear returns true if a given year is leap, and false otherwise.
func IsLeapYear(year int) bool {
	var yearIsLeap = false
	if year%4 == 0 {
		yearIsLeap = true
	}
	if year%100 == 0 && yearIsLeap == true {
		yearIsLeap = false
	}
	if year%400 == 0 {
		yearIsLeap = true
	}
	return yearIsLeap
}
