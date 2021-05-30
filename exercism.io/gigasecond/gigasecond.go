// Package gigasecond manipulates dates and durations by means of gigaseconds.
package gigasecond

// Import packages.
import "time"

// GIGASEC is the duration corresponding to one gigasecond.
const GIGASEC time.Duration = 1E9 * time.Second

// AddGigasecond adds a gigasecond to a given time specified as input.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(GIGASEC)
	return t
}
