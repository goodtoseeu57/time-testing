package main

import (
	"fmt"
	"time"
	jobs "time-testing/jobs"
)

func main() {
	duration := jobs.CalculateNextTuesdayAtNine()

	timer := time.NewTimer(duration)
	<-timer.C

	fmt.Println("run task")
}
