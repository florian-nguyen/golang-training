// Package space allows to calculate how old someone would be on any planet of our solar system given an age in seconds.
package space

// earthDaySec is the duration of one Earth year in seconds.
const earthDaySec float64 = 3600.0 * 24.0 * 365.25

// Age gives the age corresponding on a duration in seconds spent on a given planet.
func Age(t float64, p string) float64 {
	switch {
	case p == "Earth":
		return t / earthDaySec
	case p == "Mercury":
		return t / earthDaySec / 0.2408467
	case p == "Venus":
		return t / earthDaySec / 0.61519726
	case p == "Mars":
		return t / earthDaySec / 1.8808158
	case p == "Jupiter":
		return t / earthDaySec / 11.862615
	case p == "Saturn":
		return t / earthDaySec / 29.447498
	case p == "Uranus":
		return t / earthDaySec / 84.016846
	case p == "Neptune":
		return t / earthDaySec / 164.79132
	default:
		return 0.0
	}

}
