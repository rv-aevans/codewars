package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(FormatDuration(31536001))
}

func FormatDuration(seconds int64) string {
	var resSlice []string
	var resString string

	years := math.Floor(float64(seconds) / 60 / 60 / 24 / 365)
	startSeconds := seconds % (60 * 60 * 24 * 365)
	days := math.Floor(float64(startSeconds) / 60 / 60 / 24)
	startSeconds = seconds % (60 * 60 * 24)
	hours := math.Floor(float64(startSeconds) / 60 / 60)
	startSeconds = seconds % (60 * 60)
	minutes := float64(startSeconds / 60)
	finishSeconds := seconds % (60)

	if years == 1 {
		resSlice = append(resSlice, "1 year")
	} else if years > 1 {
		resSlice = append(resSlice, fmt.Sprintf("%d years", (int(years))))
	}

	if days == 1 {
		resSlice = append(resSlice, "1 day")
	} else if days > 1 {
		resSlice = append(resSlice, fmt.Sprintf("%d days", (int(days))))
	}

	if hours == 1 {
		resSlice = append(resSlice, "1 hour")
	} else if hours > 1 {
		resSlice = append(resSlice, fmt.Sprintf("%d hours", (int(hours))))
	}

	if minutes == 1 {
		resSlice = append(resSlice, "1 minute")
	} else if minutes > 1 {
		resSlice = append(resSlice, fmt.Sprintf("%d minutes", (int(minutes))))
	}

	if finishSeconds == 1 {
		resSlice = append(resSlice, "1 second")
	} else if finishSeconds > 1 {
		resSlice = append(resSlice, fmt.Sprintf("%d seconds", (int(finishSeconds))))
	}

	fmt.Println(resSlice)

	for i, val := range resSlice {
		if len(resSlice) == 1 {
			return val
		}

		if i == len(resSlice)-1 {
			return strings.TrimSuffix(resString, ", ") + " and " + val
		} else {
			resString += val + ", "
		}
	}

	return resString
}
