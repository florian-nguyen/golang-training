/*
A Pythagorean triplet is a set of three natural numbers, {a, b, c}, for
which :
a**2 + b**2 = c**2

For example :
3**2 + 4**2 = 9 + 16 = 25 = 5**2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.

Find the product a * b * c.
*/

package pythagorean

import ()

// Use this type definition,
//
//    type Triplet [3]int
//
// and implement two functions,
//
//    Range(min, max int) []Triplet
//    Sum(p int) []Triplet
//
// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
//
// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
//
// The three elements of each returned triplet must be in order,
// t[0] <= t[1] <= t[2], and the list of triplets must be in lexicographic
// order.

type Triplet [3]int

// Function triangle tells if a given set of natural numbers form a Pythagorean triplet.
func Triangle(a, b, c int) bool {
	if a*a+b*b == c*c {
		return true
	} else {
		return false
	}
}

// Function Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	var list []Triplet
	for c := min + 1; c < max+1; c++ {
		for a := min; a < c; a++ {
			for b := a; b < c; b++ {
				if Triangle(a, b, c) == true {
					list = append(list, Triplet{a, b, c})
				}
			}
		}
	}
	return list
}

// Function Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	var list []Triplet
	for c := p - 1; c > 0; c-- {
		for a := 1; a < c; a++ {
			for b := c - 1; b > a; b-- {
				if a+b+c == p && Triangle(a, b, c) {
					list = append(list, Triplet{a, b, c})
				}
			}
		}
	}
	return list
}
