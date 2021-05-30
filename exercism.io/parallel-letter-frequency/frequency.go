// Package letter calculates the frequency of characters in string lists using sequential and parallel approaches.
package letter

// Import packages.
import (
	"sync"
)

// FreqMap type stores the frequency of characters found in analyzed strings.
type FreqMap map[rune]int

// Frequency sequentially returns the frequency of each character in a string.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calculates the total frequency of each character in a slice of strings using a parallel approach (first method).
func ConcurrentFrequency(input []string) FreqMap {

	c := make(chan FreqMap)

	// Generator of FreqMap.
	for _, stringItem := range input {
		go func(item string) {
			c <- Frequency(item)
		}(stringItem)
	}

	m := FreqMap{}

	// Listener for all string values.
	for range input {
		for key, value := range <-c {
			m[key] += value
		}
	}

	return m
}
