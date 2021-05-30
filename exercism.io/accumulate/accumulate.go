package accumulate

// Accumulate returns an array of type []string where each value is the image of an input string by function f(string) string
func Accumulate(s []string, f func(string) string) (output []string) {

	output = make([]string, 0)
	for _, value := range s {
		output = append(output, f(value))
	}
	return output
}
