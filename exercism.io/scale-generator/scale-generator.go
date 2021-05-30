// Package Scale generates a music scale.
package scale

import (
	"strings"
)

var SHARP_SCALE = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var FLAT_SCALE = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var INTERVAL_SCALE = map[rune]int{'m': 1, 'M': 2, 'A': 3}

// Function Scale returns the music scale associated to the input starting note and interval.
func Scale(tonic, interval string) []string {

	// Detect which scale to use and format tonic
	sharpOrFlat := IsSharpOrFlat(tonic)
	tonic = strings.Title(tonic)

	// Rotate scale to have tonic first
	for key, value := range sharpOrFlat {
		if value == tonic {
			sharpOrFlat = append(sharpOrFlat[key:], sharpOrFlat[:key]...)
			break
		}
	}

	// Specific case if empty interval
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}

	// Compute rest of scale
	s := make([]string, 0)
	counter := 0
	for _, value := range interval {
		s = append(s, sharpOrFlat[counter])
		counter += INTERVAL_SCALE[value]
	}
	return s
}

// Function IsSharpOrFlat returns the adequate scale to be used depending on tonic parameter
func IsSharpOrFlat(tonic string) (s []string) {
	s = SHARP_SCALE
	switch tonic {
	case "d", "g", "c", "f", "bb", "eb", "F", "Bb", "Eb", "Ab", "Db", "Gb":
		s = FLAT_SCALE
	}
	return s
}
