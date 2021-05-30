package transpose

// Function Transpose computes the transpose matrix of a given string matrix
func Transpose(input []string) (t []string) {

	// Empty string case
	if len(input) == 0 {
		return []string{}
	}

	// Compute number of rows and columns (longest values)
	n_column := MaxStringLength(input)
	t = make([]string, n_column)

	// Computing transpose matrix by concatenating string characters one at a time
	for row, word := range input {
		for column, char := range word {
			t[column] += string(char)
		}
		// Space padding should only be done if a larger word exists after the current one
		for i := len(word); i < MaxStringLength(input[row:]); i++ {
			t[i] += " "
		}
	}
	return t
}

// Function maxStringLength is used to obtain the size of the largest string in a list
func MaxStringLength(s []string) (n int) {
	n = 0
	for _, word := range s {
		if len(word) > n {
			n = len(word)
		}
	}
	return n
}
