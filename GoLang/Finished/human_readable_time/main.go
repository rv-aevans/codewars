package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(humanReadable(45296))
}

func humanReadable(seconds int) string {
	if seconds > 3599 {
		return fmt.Sprintf("%02v:%02v:%02d", math.Floor(float64(seconds)/3600), seconds%3600/60, seconds%60%60%60)
	} else if seconds > 59 {
		return fmt.Sprintf("00:%02v:%02d", math.Floor(float64(seconds)/60), seconds%60)
	} else {
		return fmt.Sprintf("00:00:%02d", seconds)
	}
}
