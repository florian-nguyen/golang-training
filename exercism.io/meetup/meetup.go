package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second              = 2
	Third               = 3
	Fourth              = 4
	Last                = -1
	Teenth              = 0
)

// Day returns the meetup day
func Day(schedule WeekSchedule, weekday time.Weekday, month time.Month, year int) (output int) {

	wantedDay := make([]int, 0)
	current_date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	switch schedule {
	case First, Second, Third, Fourth, Last:
		for current_date.Month() == month {
			if current_date.Weekday() == weekday {
				wantedDay = append(wantedDay, current_date.Day())
			}
			current_date = time.Date(year, month, current_date.Day()+1, 0, 0, 0, 0, time.UTC)
		}
		if schedule != Last {
			output = wantedDay[int(schedule)-1]
			break
		} else {
			output = wantedDay[len(wantedDay)-1]
			break
		}

	case Teenth:
		// Teenth days are 13-19, so we start at 13
		current_date := time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)
		for i := 0; i < 7; i++ {
			if current_date.Weekday() == weekday {
				output = current_date.Day()
				break
			}
			current_date = time.Date(year, month, current_date.Day()+1, 0, 0, 0, 0, time.UTC)
		}

	}
	return
}
