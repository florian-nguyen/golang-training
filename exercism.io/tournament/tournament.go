// Package tournament generates a score double entry board giving information about the number of matches played, wins, losses, draws and points earned by each of the four teams competing. //
package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Defining a score map for the given tournament, and a team structure type to keep track of each team's achievements in the latter.
type scoreMap map[string]*team
type team struct {
	name                            string
	played, win, loss, draw, points int
}
type teamList []team

// Tally function counts up the wins and losses of each team. //
func Tally(reader io.Reader, writer io.Writer) error {

	scanner := bufio.NewScanner(reader)
	score := make(scoreMap)

	// Scanning input
	for scanner.Scan() {

		line := scanner.Text()

		// Comments are not to be considered
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Add information relative to each game and check for errors
		if err := score.AddGame(line); err != nil {
			return err
		}
	}

	// Get and sort teams by their scores
	teams := score.GetTeams()
	sort.Sort(teamList(teams))

	// Write output
	legends := "Team                           | MP |  W |  D |  L |  P\n"
	io.WriteString(writer, legends)
	for _, team := range teams {
		io.WriteString(writer, team.String()+"\n")
	}

	// No errror issued
	return nil
}

// addGame converts information relative to a match to a given score map. //
func (score scoreMap) AddGame(line string) error {

	words := strings.Split(line, ";")

	// Check for correct line formatting
	if len(words) != 3 {
		return fmt.Errorf("Wrong line formatting in input string: %s", line)
	}

	// See if teams are already registered in score chart
	first, firstCheck := score[words[0]]
	second, secondCheck := score[words[1]]

	// If not, add them manually
	if !firstCheck {
		first = &team{name: words[0]}
		score[words[0]] = first
	}

	if !secondCheck {
		second = &team{name: words[1]}
		score[words[1]] = second
	}

	// Depending on the outcome of the match
	switch words[2] {

	case "win":
		first.AddWin()
		second.AddLoss()

	case "loss":
		first.AddLoss()
		second.AddWin()

	case "draw":
		first.AddDraw()
		second.AddDraw()

	default:
		return fmt.Errorf("Wrong format of game outcome: %s", words[2])
	}

	// No errors returned
	return nil
}

// GetTeams converts the score in an unordered list of teams //
func (score scoreMap) GetTeams() []team {
	var teams []team
	for _, team := range score {
		teams = append(teams, *team)
	}

	return teams
}

// Len returns the number of teams in the tournament. //
func (t teamList) Len() int {
	return len(t)
}

// Swap switches the order of two teams in a list. //
func (t teamList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Less returns true if i should be ranked better than j. The name refers to the "sort" package. //
func (t teamList) Less(i, j int) bool {

	// Compare points if not equal, and number of victories otherwise
	if t[i].points != t[j].points {
		return t[i].points > t[j].points
	} else if t[i].win != t[j].win {
		return t[i].win > t[j].win
	}

	// Compare names otherwise
	return t[i].name < t[j].name
}

// AddWin adds a win to the team. //
func (t *team) AddWin() {
	t.played++
	t.win++
	t.points += 3
}

// AddLoss adds a loss to the team. //
func (t *team) AddLoss() {
	t.played++
	t.loss++
}

// AddDraw adds a draw to the team. //
func (t *team) AddDraw() {
	t.played++
	t.draw++
	t.points++
}

// String returns the string associated to each team's name, points, matches played, etc. //
func (t *team) String() string {
	fmtString := "%-31s| %2d | %2d | %2d | %2d | %2d"
	return fmt.Sprintf(fmtString, t.name, t.played, t.win, t.draw, t.loss, t.points)
}
