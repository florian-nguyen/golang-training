package main

import (
	"io"
	"log"
	"os"
	"sync"
	"time"

	"gitlab.com/florian-nguyen/training/nhl-fetch-players/nhlApi"
)

func main() {

	// To help benchmarking the request time
	now := time.Now()

	rosterFile, err := os.OpenFile("rosters.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Error opening the file rosters.txt: %v", err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	teams, err := nhlApi.GetAllTeams()
	if err != nil {
		log.Fatalf("Error while getting all teams: %v", err)
	}

	var wg sync.WaitGroup

	// Setting size of WaitGroup : the latter will wait until that counter reaches 0
	wg.Add(len(teams))

	// Unbuffered channnel
	results := make(chan []nhlApi.Roster)

	for _, team := range teams {
		// log.Println("--------------------------")
		// log.Printf("Name: %s", team.Name)
		// log.Println("--------------------------")

		go func(team nhlApi.Team) {
			roster, err := nhlApi.GetRosters(team.ID)
			if err != nil {
				log.Fatalf("Error getting roster: %v", err)
			}

			// Sending result to the channel
			results <- roster

			// Decrement WaitGroup by 1
			wg.Done()

		}(team)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	display(results)

	log.Printf("took %v", time.Now().Sub(now).String())
}

func display(results chan []nhlApi.Roster) {
	for r := range results {
		for _, ros := range r {
			log.Println("--------------------------")
			log.Printf("ID: %d\n", ros.Person.ID)
			log.Printf("Name: %s\n", ros.Person.Fullname)
			log.Printf("Position: %s\n", ros.Position.Abbreviation)
			log.Printf("Jersey: %s\n", ros.Jerseynumber)
			log.Println("--------------------------")
		}
	}
}
