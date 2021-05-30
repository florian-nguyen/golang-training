// Package clock handles times without dates, i.e. in the range of 00:00 to 23:59
package clock

// Import packages.
import (
	"fmt"
)

// Clock type is defined with an hour and a number of minutes.
type Clock int

// New generates a new clock.
func New(h, m int) Clock {
	var minCount int = (h*60 + m) % (24 * 60)
	if minCount < 0 {
		minCount += 24 * 60
	}
	return Clock(minCount)
}

// String will return time.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c-(c/60)*60)
}

// Add will add a given amount of minutes to the hour.
func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}
