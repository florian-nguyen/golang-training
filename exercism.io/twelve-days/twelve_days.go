// Package twelve gives the verses of the 'The Twelve Days of Christmas' song.
package twelve

import (
	"fmt"
	"strings"
)

// Beginning of each verse.
const start = "On the %s day of Christmas my true love gave to me"

// Items obtained on each new day.
var presents = [][]string{
	[]string{"first", "a Partridge in a Pear Tree."},
	[]string{"second", "two Turtle Doves"},
	[]string{"third", "three French Hens"},
	[]string{"fourth", "four Calling Birds"},
	[]string{"fifth", "five Gold Rings"},
	[]string{"sixth", "six Geese-a-Laying"},
	[]string{"seventh", "seven Swans-a-Swimming"},
	[]string{"eighth", "eight Maids-a-Milking"},
	[]string{"ninth", "nine Ladies Dancing"},
	[]string{"tenth", "ten Lords-a-Leaping"},
	[]string{"eleventh", "eleven Pipers Piping"},
	[]string{"twelfth", "twelve Drummers Drumming"},
}

// Song simply returns the actual song lyrics.
func Song() string {
	var song string
	for i := 0; i < 12; i++ {
		song = song + Verse(i+1) + "\n"
	}
	return song
}

// Verse returns the lyrics associated to a given verse of the song.
func Verse(n int) string {
	lyricsList := []string{fmt.Sprintf(start, presents[n-1][0])}
	for i := n - 1; i >= 0; i-- {
		if i == 0 && n != 1 {
			lyricsList = append(lyricsList, "and "+presents[i][1])
		} else {
			lyricsList = append(lyricsList, presents[i][1])
		}
	}
	return strings.Join(lyricsList, ", ")
}
