package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

/*
There are 12 children in the class:
- Alice, Bob, Charlie, David,
- Eve, Fred, Ginny, Harriet,
- Ileana, Joseph, Kincaid, and Larry.
*/

var PLANT_TABLE = map[rune]string{
	'V': "violets",
	'R': "radishes",
	'G': "grass",
	'C': "clover",
}

type Garden struct {
	plan map[string][]string
}

func NewGarden(diagram string, children []string) (*Garden, error) {

	g := new(Garden)
	lines := strings.Split(diagram, "\n")[1:]

	// Exit if incorrect input format
	if len(lines) != 2 || len(lines[0]) != 2*len(children) || len(lines[1]) != 2*len(children) || strings.ToUpper(diagram) != diagram {
		return nil, errors.New("Error: Input diagram format is incorrect!")
	}

	// Initialize map and exit if duplicates detected
	// Warning: when using sort, be careful to work on a copy of the children array, as function arguments should not be modified
	g.plan = map[string][]string{}
	namelist := make([]string, len(children))
	copy(namelist, children)
	sort.Strings(namelist)
	for _, name := range children {
		if _, ok := g.plan[name]; ok {
			return nil, errors.New("Error: Name duplicate found!")
		}
		g.plan[name] = make([]string, 0)
	}

	// Read lines and assign plants to children
	for _, line := range lines {
		// fmt.Println(line)
		for key, name := range namelist {
			g.plan[name] = append(g.plan[name], PLANT_TABLE[rune(line[2*key])])
			g.plan[name] = append(g.plan[name], PLANT_TABLE[rune(line[2*key+1])])
		}
	}
	return g, nil
}

func (g *Garden) Plants(name string) ([]string, bool) {
	if _, ok := g.plan[name]; !ok {
		return []string{}, false
	}
	return g.plan[name], true
}
