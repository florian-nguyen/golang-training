/*
Given a number, find the sum of all the unique multiples of particular numbers up to
but not including that number.

If we list all the natural numbers below 20 that are multiples of 3 or 5,
we get 3, 5, 6, 9, 10, 12, 15, and 18.

The sum of these multiples is 78.
*/

// Package summultiples returns the sum of all integers up to a given limit that are also divisable by the also specified integers.
package summultiples

import ()

// Returns the sum of all the unique multiples of particular numbers up to but not including that number.
func SumMultiples(n int, divisors ...int) int {
	list := make([]int, n)
	var counter int = 0
	for i := 1; i < n; i++ {
		list[i] = 0
	}
	for _, div := range divisors {
		if div != 0 {
			for i := 1; i < n; i++ {
				if i%div == 0 && div != 0 {
					list[i] = i
				}
			}
		}
	}
	for i := 1; i < n; i++ {
		counter += list[i]
	}
	return counter
}
