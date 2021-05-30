/*
The Collatz Conjecture or 3x+1 problem can be summarized as follows:

Take any positive integer n. If n is even, divide n by 2 to get n / 2. If n is
odd, multiply n by 3 and add 1 to get 3n + 1. Repeat the process indefinitely.
The conjecture states that no matter which number you start with, you will
always reach 1 eventually.

Given a number n, return the number of steps required to reach 1.
*/

// Package collatz allows the use of the Collatz Conjecture and gives the required number of steps required to achieve 1 for a given integer number n.
package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	var n_iter, n_next int
	n_next = n
	n_iter = 0

	if n_next <= 0 {
		return 0, errors.New("Error: Input must be a strictly positive integer!")
	}
	for n_next != 1 {
		switch n_next % 2 {
		case 0:
			n_next = n_next / 2
			n_iter++
		case 1:
			n_next = 3*n_next + 1
			n_iter++
		}
	}

	return n_iter, nil

}
