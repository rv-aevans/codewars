package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Month struct {
	Month string
	Temp  float64
}

type City struct {
	City   string
	Months []Month
}

func main() {

	inputText := "Rome:Jan 81.2,Feb 63.2,Mar 70.3,Apr 55.7,May 53.0,Jun 36.4,Jul 17.5,Aug 27.5,Sep 60.9,Oct 117.7,Nov 111.0,Dec 97.9\nLondon:Jan 48.0,Feb 38.9,Mar 39.9,Apr 42.2,May 47.3,Jun 52.1,Jul 59.5,Aug 57.2,Sep 55.4,Oct 62.0,Nov 59.0,Dec 52.9"

	var cityStructSlice []City

	var citySlice []string
	citySlice = strings.Split(inputText, "\n")
	for _, city := range citySlice {
		newCity := City{}
		nameSplit := strings.Split(city, ":")
		newCity.City = nameSplit[0]
		monthSplit := strings.Split(nameSplit[1], ",")
		for _, month := range monthSplit {
			splitMonth := strings.Split(month, " ")
			n, _ := strconv.ParseFloat(splitMonth[1], 64)
			newCity.Months = append(newCity.Months,
				Month{
					Month: splitMonth[0],
					Temp:  n,
				},
			)
		}
		cityStructSlice = append(cityStructSlice, newCity)
	}

	for _, city := range cityStructSlice {
		total := 0
		for _, month := range city.Months {
			total = total + int(month.Temp)
		}
		result := total / 12
		fmt.Println(city, result)
	}
}

func Solution() {
	// Split the string by ","
}
