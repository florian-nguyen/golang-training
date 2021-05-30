// Package twofer simply returns a string that indicates that something is shared between two people.
package twofer

// ShareWith shares something between me and someone whose name is given as input string.
func ShareWith(input string) string {

	// Depending on the nature of the specified name...
	switch {
	case input != "you" && input != "":
		return "One for " + input + ", one for me." // General case.
	default:
		return "One for you, one for me." // if the input is "me" or not a name.
	}
}
