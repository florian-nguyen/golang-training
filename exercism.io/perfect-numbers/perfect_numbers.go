package perfect

import (
	"errors"
)

// Classification type defined here
type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

// ErrOnlyPositive is to be returned if the input is not a positive integer
var ErrOnlyPositive = errors.New("Error: Input can only be a positive number!")

// Classify returns integer n's classification based on its Aliquot Sum
func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	switch {
	case AliquotSum(n) == n:
		return ClassificationPerfect, nil
	case AliquotSum(n) < n:
		return ClassificationDeficient, nil
	case AliquotSum(n) > n:
		return ClassificationAbundant, nil
	default:
		return 0, ErrOnlyPositive
	}
}

// AliquotSum implements the Aliquot Sum algorithm
func AliquotSum(n int64) (output int64) {

	factors := make([]int64, 0)

	// Obtain all factors
	for i := int64(1); i < n/2+1; i++ {
		if n%int64(i) == 0 {
			factors = append(factors, i)
		}
	}

	// Compute sum
	for i := 0; i < len(factors); i++ {
		output += factors[i]
	}
	return output
}
