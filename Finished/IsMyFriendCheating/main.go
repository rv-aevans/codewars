package main

import (
	"fmt"
	"math"
	"time"
)

type tenure struct {
	Years  int
	Months int
}

func main() {
	hireDate := time.Date(2015, 11, 14, 10, 45, 16, 0, time.UTC)

	days := time.Now().Sub(hireDate).Hours() / 24
	years := math.Floor(days / 365)
	monthsRemainder := int(math.Floor(days/30) - (12 * years))

	fmt.Println(
		tenure{
			Years:  int(years),
			Months: monthsRemainder,
		},
	)
}
