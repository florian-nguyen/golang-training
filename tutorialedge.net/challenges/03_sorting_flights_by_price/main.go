package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {

	var m = map[int]Flight{}
	for _, flight := range flights {
		m[flight.Price] = flight
	}

	prices := []int{}
	for key := range m {
		prices = append(prices, key)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(prices)))

	orderedFlights := []Flight{}
	for _, price := range prices {
		orderedFlights = append(orderedFlights, Flight{
			Origin:      m[price].Origin,
			Destination: m[price].Destination,
			Price:       price,
		})
	}

	return orderedFlights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d\n", flight.Origin, flight.Destination, flight.Price)
	}
}

func main() {
	// an empty slice of flights
	var flights []Flight

	origins := []string{"Paris", "New York", "Tokyo", "Los Angeles", "Osaka"}
	destinations := []string{"Madrid", "London", "Vancouver", "Toronto", "Toulouse"}
	for key := range origins {
		flights = append(flights, Flight{
			Origin:      origins[key],
			Destination: destinations[key],
			Price:       rand.Intn(1000) + 1000,
		})
	}

	sortedList := SortByPrice(flights)
	printFlights(sortedList)
}

/* Better solution :

package main

import (
	"fmt"
	"sort"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type ByPrice []Flight

func (p ByPrice) Len() int {
	return len(p)
}

func (p ByPrice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByPrice) Less(i, j int) bool {
	return p[i].Price > p[j].Price
}

func SortByPrice(flights []Flight) []Flight {
	sort.Sort(ByPrice(flights))
	return flights
}

func main() {
	flights := []Flight{
		Flight{Price: 30},
		Flight{Price: 20},
		Flight{Price: 50},
		Flight{Price: 1000},
	}

	sort.Sort(ByPrice(flights))

	fmt.Println(flights)
}

*/
