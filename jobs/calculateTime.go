package jobs

import "time"

var now = time.Now()

func CalculateNextTuesdayAtNine() time.Duration {
	location := now.Location()
	daysUntilTuesday := int(time.Tuesday - now.Weekday())

	if daysUntilTuesday < 0 || (daysUntilTuesday == 0 && now.Hour() >= 9) {
		daysUntilTuesday += 7
	}

	nextTuesdayAtNine := time.Date(now.Year(), now.Month(), now.Day()+daysUntilTuesday, 9, 0, 0, 0, location)

	return nextTuesdayAtNine.Sub(now)
}
