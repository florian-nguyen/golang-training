// Package wordcount counts the number of occurrences of each word in a given string.
package wordcount

import (
	"regexp"
	//"sort"
	"strings"
)

// Type Frequency is defined to store the number of occurrences of each word.
type Frequency map[string]int

// Function WordCount returns the Frequency map associated to a given input phrase in string format.
/*func WordCount(phrase string) Frequency {

	// Convert all possible separators into spaces, and convert to lowercase
	s := strings.ToLower(strings.Replace(phrase, "\n", " ", -1))
	s = strings.Replace(s, ",", " ", -1)
	s = strings.TrimSpace(s)

	// Remove all ponctuation characters
	s = strings.Replace(s, "!", "", -1)
	s = strings.Replace(s, "&", "", -1)
	s = strings.Replace(s, "@", "", -1)
	s = strings.Replace(s, "%", "", -1)
	s = strings.Replace(s, "$", "", -1)
	s = strings.Replace(s, "%", "", -1)
	s = strings.Replace(s, "^", "", -1)
	s = strings.Replace(s, ":", "", -1)
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, ":", "", -1)

	// Treat apostrophs as a separate case
	s = strings.Replace(s, "n't", "$", -1)
	s = strings.Replace(s, "'", "", -1)
	s = strings.Replace(s, "$", "n't", -1)

	// Remove multiple spaces

	for strings.Contains(s, "  ") {
		s = strings.Replace(s, "  ", " ", -1)
	}

	var list []string = strings.Split(s, " ")
	list = sort.StringSlice(list)
	var freq Frequency = make(Frequency)
	for _, value := range list {
		freq[value]++
	}
	return freq
}*/

func WordCount(phrase string) Frequency {

	count := make(Frequency)

	r := regexp.MustCompile(`[\w']+`)
	words := r.FindAllString(strings.ToLower(phrase), -1)

	for _, word := range words {
		word = strings.Trim(word, "'")
		count[word]++
	}

	return count
}
