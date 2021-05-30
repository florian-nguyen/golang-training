package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	seen := map[string]Developer{}
	for _, d := range developers {
		_, ok := seen[d.Name]
		if !ok {
			seen[d.Name] = d
		}
	}

	filteredList := []string{}
	for _, value := range seen {
		filteredList = append(filteredList, value.Name)
	}
	return filteredList
}

func main() {
	fmt.Println("Filter Unique Challenge")
}
