package strain

// Ints type definition
type Ints []int

// Lists type definition
type Lists [][]int

// Strings type definition
type Strings []string

// Keep if function is true
func (n Ints) Keep(f func(int) bool) (output Ints) {

	// Exit if empty list
	if n == nil {
		return nil
	}

	for _, value := range n {
		if f(value) == true {
			output = append(output, value)
		}
	}
	return output
}

// Discard if function is false
func (n Ints) Discard(f func(int) bool) (output Ints) {

	// Exit if empty list
	if n == nil {
		return nil
	}

	for _, value := range n {
		if f(value) == false {
			output = append(output, value)
		}
	}
	return output
}

// Keep if function is true
func (l Lists) Keep(f func([]int) bool) (output Lists) {

	// Exit if empty list
	if l == nil {
		return nil
	}

	for _, value := range l {
		if f(value) == true {
			output = append(output, value)
		}
	}
	return output
}

// Keep if function is true
func (s Strings) Keep(f func(string) bool) (output Strings) {

	// Exit if empty list
	if s == nil {
		return nil
	}

	for _, value := range s {
		if f(value) == true {
			output = append(output, value)
		}
	}
	return output
}
