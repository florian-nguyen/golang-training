// Package spiralmatrix allows the creation of spiral matrices, for instance :
//  1  2  3 4
// 12 13 14 5
// 11 16 15 6
// 10  9  8 7
package spiralmatrix

import ()

// Function SpiralMatrix generates a spiral matrix of size n.
func SpiralMatrix(n int) [][]int {
	var counter int = 1
	var matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < (n+1)/2; i++ {
		// From UP-LEFT to UP-RIGHT
		for k := i; k < n-i; k++ {
			matrix[i][k] = counter
			counter++
		}
		// From UP-RIGHT to DOWN-RIGHT
		for j := i + 1; j < n-i; j++ {
			matrix[j][n-1-i] = counter
			counter++
		}
		// From DOWN-RIGHT to DOWN-LEFT
		for k := n - i - 2; k >= i; k-- {
			matrix[n-1-i][k] = counter
			counter++
		}
		// From DOWN-LEFT to UP-LEFT
		for j := n - i - 2; j >= i+1; j-- {
			matrix[j][i] = counter
			counter++
		}
	}
	return matrix
}
