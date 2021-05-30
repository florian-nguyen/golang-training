package prime

// Factors returns the prime factors of a given number
func Factors(input int64) []int64 {

	factors := make([]int64, 0)
	current := int64(2)

	for input != int64(1) {
		for input%current == 0 {
			input /= current
			factors = append(factors, current)
		}
		current++
	}

	return factors
}
