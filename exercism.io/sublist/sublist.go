// Package sublist indicates if the first list is a sublist or a superlist of the second one. The first item is a superlist if it contains the second list, a sublist if it is contained in the latter.
package sublist

// Examples
// * A = [1, 2, 3], B = [1, 2, 3, 4, 5], A is a sublist of B
// * A = [3, 4, 5], B = [1, 2, 3, 4, 5], A is a sublist of B
// * A = [3, 4], B = [1, 2, 3, 4, 5], A is a sublist of B
// * A = [1, 2, 3], B = [1, 2, 3], A is equal to B
// * A = [1, 2, 3, 4, 5], B = [2, 3, 4], A is a superlist of B
// * A = [1, 2, 4], B = [1, 2, 3, 4, 5], A is not a superlist of, sublist of or equal to B

type Relation string

const (
	equal     Relation = "equal"
	sublist   Relation = "sublist"
	superlist Relation = "superlist"
	unequal   Relation = "unequal"
)

// Function isEqual compares two int[] with equal length values
func IsEqual(a, b []int) Relation {
	var counter int
	counter = 0
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			counter++
		}
	}
	if counter == len(a) {
		return equal
	}
	return unequal

}

// Function sublist returns a string indicating whether the first list is a superlist/sublist of the second one, or none of these.
func Sublist(a, b []int) Relation {

	var counter int = 0

	// Same length case
	if len(a) == 0 && len(b) != 0 {
		return sublist

	} else if len(b) == 0 && len(a) != 0 {
		return superlist

	} else if len(a) == len(b) {
		return IsEqual(a, b)

	} else if len(a) < len(b) {
		for i0 := 0; i0 < len(b)-len(a)+1; i0++ {
			if b[i0] == a[0] {
				counter = 0
				for j := 0; j < len(a); j++ {
					if b[i0+j] == a[j] {
						counter++
					}
				}
				if counter == len(a) {
					return sublist
				}
			}
		}

	} else if len(a) > len(b) {
		for i0 := 0; i0 < len(a)-len(b)+1; i0++ {
			if a[i0] == b[0] {
				counter = 0
				for j := 0; j < len(b); j++ {
					if a[i0+j] == b[j] {
						counter++
					}
				}
				if counter == len(b) {
					return superlist
				}
			}
		}
	}

	// Return unequal if none of the above creteria are met
	return unequal

}

// Other alternative
/*
func testSublist(list1, list2 []int) bool {
	for i2 := range list2 {
		match := true
		for i1, ch1 := range list1 {
			idx2 := i2 + i1
			if len(list2) <= idx2 || list2[idx2] != ch1 {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return len(list1) == 0
}

func Sublist(list1, list2 []int) Relation {
	sublist := testSublist(list1, list2)
	superlist := testSublist(list2, list1)
	switch {
	case sublist && superlist:
		return "equal"
	case sublist:
		return "sublist"
	case superlist:
		return "superlist"
	default:
		return "unequal"
	}
}
*/
