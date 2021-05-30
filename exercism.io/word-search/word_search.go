package wordsearch

import (
	"errors"
)

/* BETTER SOLUTION
var bearings = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func Solve(words, puzzle []string) (map[string][2][2]int, error) {
	location := make(map[string][2][2]int)
	for _, w := range words {
		for i, row := range puzzle {
			for j := 0; j < len(row); j++ {
			NextBearing:
				for _, b := range bearings {
					ii, jj := i, j
					for cidx := 0; cidx < len(w); cidx, ii, jj = cidx+1, ii+b[0], jj+b[1] {
						if ii < 0 || jj < 0 || ii >= len(puzzle) || jj >= len(row) || puzzle[ii][jj] != w[cidx] {
							continue NextBearing
						}
					}
					location[w] = [2][2]int{{j, i}, {jj - b[1], ii - b[0]}}
				}
			}
		}
		if _, ok := location[w]; !ok {
			return nil, errors.New("No match")
		}
	}
	return location, nil
*/

// Solve returns, for each word to find in a puzzle, the position of the first and last letter of the word in the puzzle
func Solve(words, puzzle []string) (map[string][2][2]int, error) {

	m := make(map[string][2][2]int, 0)

	// For each letter of the word to find, search the puzzle...
	for _, word := range words {

		var match_found bool
		for row, line := range puzzle {
			for col, char := range line {

				if string(char) == string(word[0]) {

					if result, ok := SearchRow(puzzle, row, col, word); ok && !match_found {
						m[word] = result
						match_found = true
					}

					if result, ok := SearchCol(puzzle, row, col, word); ok && !match_found {
						m[word] = result
						match_found = true
					}

					if result, ok := SearchDiag(puzzle, row, col, word); ok && !match_found {
						m[word] = result
						match_found = true
					}
				}
			}
		}

		if !match_found {
			return nil, errors.New("Error: No match found for word '" + word + "'!")
		}
	}
	return m, nil
}

// SearchRow searches for a match for a given word in a row where the first letter of that word has been found at a certain index
func SearchRow(s []string, row, col int, word string) ([2][2]int, bool) {

	// Search the current row in the left-to-right order
	if len(s[row][col:]) >= len(word) && s[row][col:col+len(word)] == word {
		return [2][2]int{{col, row}, {col + len(word) - 1, row}}, true
	}

	// Same in the right-to-left order
	if len(s[row][:col+1]) >= len(word) && ReverseString(s[row][col-len(word)+1:col+1]) == word {
		return [2][2]int{{col, row}, {col - len(word) + 1, row}}, true
	}

	return [2][2]int{}, false
}

// SearchCol searches for a match for a given word in a column where the first letter of that word has been found at a certain index
func SearchCol(s []string, row, col int, word string) ([2][2]int, bool) {

	// Compute the column array
	var c string
	for i := 0; i < len(s); i++ {
		c += string(s[i][col])
	}

	// Search the current column in the top-to-bottom order
	if len(c[row:]) >= len(word) && c[row:row+len(word)] == word {
		return [2][2]int{{col, row}, {col, row + len(word) - 1}}, true
	}

	// Search the current column in the top-to-bottom order
	if len(c[:row+1]) >= len(word) && ReverseString(c[row-len(word)+1:row+1]) == word {
		return [2][2]int{{col, row}, {col, row - len(word) + 1}}, true
	}

	return [2][2]int{}, false
}

// SearchDiag searches for a match for a given word in the diagonals intersecting where a the first letter of that word has been found
func SearchDiag(s []string, row, col int, word string) ([2][2]int, bool) {

	// Compute diagonal strings
	var d1, d2 string
	var i1, i2 int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[0]); j++ {
			// '\' direction diagonal
			if i-j == row-col {
				d1 += string(s[i][j])
				if i == row {
					i1 = len(d1) - 1
				}
			}

			// '/' direction diagonal
			if i+j == row+col {
				d2 += string(s[i][j])
				if i == row {
					i2 = len(d2) - 1
				}
			}
		}
	}

	// Search diagonals in the current/reverse order
	if len(d1[i1:]) >= len(word) && d1[i1:i1+len(word)] == word {
		return [2][2]int{{col, row}, {col + len(word) - 1, row + len(word) - 1}}, true
	}
	if len(d2[i2:]) >= len(word) && d2[i2:i2+len(word)] == word {
		return [2][2]int{{col, row}, {col - len(word) + 1, row + len(word) - 1}}, true
	}

	if len(d1[:i1+1]) >= len(word) && ReverseString(d1[i1-len(word)+1:i1+1]) == word {
		return [2][2]int{{col, row}, {col - len(word) + 1, row - len(word) + 1}}, true
	}
	if len(d2[:i2+1]) >= len(word) && ReverseString(d2[i2-len(word)+1:i2+1]) == word {
		return [2][2]int{{col, row}, {col + len(word) - 1, row - len(word) + 1}}, true
	}
	// fmt.Println(d1, d2, word)
	return [2][2]int{}, false
}

// ReverseString returns the runes of a string
func ReverseString(s string) string {
	var output string
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		output += string(runes[len(runes)-1-i])
	}
	return output
}
