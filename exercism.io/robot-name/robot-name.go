// Package robotname generates robots with random names.
package robotname

// Import packages.
import (
	"fmt"
	"math/rand"
	"time"
)

// letters stores all possible letters for name attribution.
const letters = "ABCDEFGHIJKLMNOPQRSTUVWXZ"

// Robot type is characterized by a name.
type Robot string

// inAction is a list of robots currently in action.
var inAction = make(map[Robot]bool)

// Name renames the robot by a non-existing value.
func (r *Robot) Name() string {
	// Generate new seed.
	rand.Seed(time.Now().UnixNano())
	// for key, value := string(*r) == "", true; value && key; key, value = inAction[*r] {
	if string(*r) == "" {
		*r = Robot(letter() + letter() + number())
	}
	inAction[*r] = true
	return string(*r)
}

// letter returns a random letter within 'A'-'Z'.
func letter() string {
	return string(letters[rand.Intn(len(letters))])
}

// number returns a random integer within 0 and 999 (three digits used).
func number() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}

// Reset the robot's name.
func (r *Robot) Reset() {
	*r = Robot("")
}
