// Package diffsquares calculates the square of a sum, the sum of squares and their difference.
package diffsquares

// Difference calculates the difference between the square of the sum and the sum of the squares of the n first integers.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

// SumOfSquares calculates the sum of the n first squared integers.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// SquareOfSums calculates the square of the sum of the n first integers.
func SquareOfSums(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}
