// Package Sieve identifies all the Prime numbers from 2 up to a given limit using the Sieve of Eratosthenes.
package sieve

// Function Sieve returns all the prime numbers between 2 and a given limit using the Sieve of Erastosthenes.
func Sieve(limit int) []int {

	// Initialize array with all integers of list
	prime := make([]int, 0)
	list := make(map[int]bool)
	for i := 2; i < limit+1; i++ {
		list[i] = true
	}

	// Sieve of Erastosthenes
	for i := 2; i < limit+1; i++ {
		if list[i] == true {
			prime = append(prime, i)
			for j := i + 1; j < limit+1; j++ {
				if list[j] == true && j%i == 0 {
					list[j] = false
				}
			}
		}
	}

	return prime
}
