package listops

type IntList []int
type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Filter removes from a list all the values for which the argument function is false
func (l IntList) Filter(fn predFunc) IntList {
	output := make(IntList, 0)
	for _, value := range l {
		if fn(value) == true {
			output = append(output, value)
		}
	}
	return output
}

// Length returns the length of the list
func (l IntList) Length() int {
	return len(l)
}

// Foldr applies a binary function to an integer list from the right
func (l IntList) Foldr(fn binFunc, n int) int {
	if len(l) == 0 {
		return n
	}
	previous := n
	for i := len(l) - 1; i > -1; i-- {
		previous = fn(l[i], previous)
	}
	return previous
}

// Foldr applies a binary function to an integer list from the left
func (l IntList) Foldl(fn binFunc, n int) int {
	if len(l) == 0 {
		return n
	}
	previous := n
	for _, value := range l {
		previous = fn(previous, value)
	}
	return previous
}

// Map applies a function to each integer of a list
func (l IntList) Map(fn unaryFunc) IntList {
	output := make(IntList, 0)
	for _, value := range l {
		output = append(output, fn(value))
	}
	return output
}

// Reverse returns a new IntList in the opposite order
func (l IntList) Reverse() IntList {
	output := make(IntList, 0)
	for i := 0; i < len(l); i++ {
		output = append(output, l[len(l)-1-i])
	}
	return output
}

// Append appends an IntList to another
func (l IntList) Append(newList IntList) IntList {
	output := make(IntList, 0)
	for _, value := range l {
		output = append(output, value)
	}
	for _, value := range newList {
		output = append(output, value)
	}
	return output
}

// Concat concatenates several IntList together
func (l IntList) Concat(input []IntList) IntList {
	output := make(IntList, 0)
	output = output.Append(l)
	for _, value := range input {
		output = output.Append(value)
	}
	return output
}
