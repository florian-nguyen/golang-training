// Package pascal computes Pascal's triangle up to a given number of rows.
package pascal

import ()

// Function Triangle returns the triangle up to the given integer number in
// the form of a [][]int.
func Triangle(n int) [][]int {

	var triangle = make([][]int, n)
	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				triangle[i][j] = 1
			} else {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}
	}
	return triangle

}
