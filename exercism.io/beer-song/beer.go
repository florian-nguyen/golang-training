// Package beer gives the lyrics of the well-known beer song.
package beer

import (
	"errors"
	"strconv"
)

// Function Verse returns the nth verse of the song.
func Verse(n int) (string, error) {
	// Error cases
	if n < 0 || n > 99 {
		return "", errors.New("Error: Verse index must be a positive integer between 0 and 99.\n")
	}

	// General cases
	switch n {
	case 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case 1:
		return strconv.Itoa(n) + " bottle of beer on the wall, " + strconv.Itoa(n) + " bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case 2:
		return strconv.Itoa(n) + " bottles of beer on the wall, " + strconv.Itoa(n) + " bottles of beer.\nTake one down and pass it around, " + strconv.Itoa(n-1) + " bottle of beer on the wall.\n", nil
	default:
		return strconv.Itoa(n) + " bottles of beer on the wall, " + strconv.Itoa(n) + " bottles of beer.\nTake one down and pass it around, " + strconv.Itoa(n-1) + " bottles of beer on the wall.\n", nil
	}
}

// Function Verses returns the verses between two specified indexes start and stop.
func Verses(start, stop int) (string, error) {
	output := ""
	for i := start; i >= stop; i-- {
		if verse, err := Verse(i); err == nil {
			output += verse + "\n"
		} else {
			return "", err
		}
	}
	if output == "" {
		return "", errors.New("Error: Empty verses detected.")
	}
	return output, nil
}

// Function Song returns the complete beer song.
func Song() string {
	if verses, err := Verses(99, 0); err == nil {
		return verses
	}
	return ""
}
