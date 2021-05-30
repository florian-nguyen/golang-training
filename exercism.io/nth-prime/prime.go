/*
Given a number n, determine what the nth prime is.

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that
the 6th prime is 13.

If your language provides methods in the standard library to deal with prime
numbers, pretend they don't exist and implement them yourself.
*/

// Package prime allows the determination of the Nth Prime number.
package prime

// Returns the value of the Nth prime number.
func Nth(n int) (int, bool) {
	var primes []int = []int{2}
	var current int = 2
	var currentIsNotPrime bool
	if n <= 0 {
		return 0, false
	} else if n == 1 {
		return primes[n-1], true
	} else {
		for len(primes) < n {
			currentIsNotPrime = false
			for _, value := range primes {
				if current%value == 0 {
					currentIsNotPrime = true
					break
				}
			}
			if !currentIsNotPrime {
				primes = append(primes, current)
			}
			current++
		}
		return primes[n-1], true
	}
	return 0, false
}
