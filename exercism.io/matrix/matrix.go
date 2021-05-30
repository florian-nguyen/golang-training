// Package matrix gives the rows and columns of a matrix of variable size.
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Type matrix definition
type matrix struct {
	rows, cols int
	coef       []int
}

// New generates a new matrix.
func New(s string) (*matrix, error) {

	// Reading string and obtaining matrix size
	lines := strings.Split(s, "\n")
	rowSize := len(lines)
	colSize := len(strings.Split(lines[0], " "))
	matrix := matrix{rows: rowSize, cols: colSize}

	// Deducing matrix coefficients
	for _, row := range lines {

		// Remove spaces and compute coefficients by row
		rowCoef := strings.Split(strings.TrimSpace(row), " ")
		if matrix.cols != len(rowCoef) {
			return nil, errors.New("Error: All columns must be of same length!")
		}

		// Convert strings to integer and deal with errors
		for _, rowChar := range rowCoef {
			coef, err := strconv.Atoi(rowChar)

			if err != nil {
				return nil, err
			}

			// Assign to matrix
			matrix.coef = append(matrix.coef, coef)
		}
	}

	return &matrix, nil
}

// Rows returns the rows of a matrix.
func (m matrix) Rows() [][]int {
	rows := make([][]int, m.rows)
	for r := 0; r < m.rows; r++ {
		rows[r] = make([]int, m.cols)
		for c := 0; c < m.cols; c++ {
			rows[r][c] = m.coef[r*m.cols+c]
		}
	}
	return rows
}

// Cols returns the columns of a matrix.
func (m matrix) Cols() [][]int {
	cols := make([][]int, m.cols)
	for c := 0; c < m.cols; c++ {
		cols[c] = make([]int, m.rows)
		for r := 0; r < m.rows; r++ {
			cols[c][r] = m.coef[r*m.cols+c]
		}
	}
	return cols
}

// Set allows to change a coefficient of a matrix.
func (m matrix) Set(r, c, value int) bool {
	// Check the specified row and column indexes.
	if !(r < m.rows && r > -1) {
		return false
	}
	if !(c < m.cols && c > -1) {
		return false
	}

	// Proceed to coefficient change
	m.coef[r*m.cols+c] = value
	return true
}
