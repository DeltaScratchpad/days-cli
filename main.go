package main

import (
	"fmt"
	"math"
	"time"

	"flag"
)

func main() {
	yearNow, monthNow, _ := time.Now().Date()

	year := flag.Int("year", yearNow, "What year?")
	month := flag.Int("month", int(monthNow), "What month?")
	ratio := flag.Float64("ratio", 0.2, "How much of the month do you need to be in the office?")
	daysOff := flag.Int("off", 0, "How many days off did you take?")

	// set custom usage
	flag.Usage = func() {
		fmt.Print("Calculate hybrid days needed (rounded to nearest 0.5).\n")
		fmt.Printf("Flags: -year  %d, -month %d, -ratio  0.2, -off 0\n\n", yearNow, int(monthNow))
	}

	flag.Parse()
	// if no flags provided, show usage and use defaults
	if flag.NFlag() == 0 {
		flag.Usage()
	}

	// Get number of weekdays in the month

	firstDay := time.Date(*year, time.Month(*month), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1)
	daysInMonth := lastDay.Day()

	weekdays := 0
	for day := 1; day <= daysInMonth; day++ {
		currentDay := time.Date(*year, time.Month(*month), day, 0, 0, 0, 0, time.UTC)
		weekday := currentDay.Weekday()

		if weekday >= time.Monday && weekday <= time.Friday {
			weekdays++
		}
	}
	fmt.Printf("Year: %d, Month: %d has %d weekdays\n", *year, *month, weekdays)
	worked := weekdays - *daysOff

	owed := float64(worked) * *ratio
	rounded := math.Ceil(owed*2) / 2

	fmt.Printf("You need to be in the office: %g days", rounded)

}
