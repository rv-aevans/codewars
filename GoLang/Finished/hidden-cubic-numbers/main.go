package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	testString = "aqdf&0#1xyz!22[153(777.777"
)

func main() {
	fmt.Println(IsSumOfCubes("aqdf&0#1xyz!22[153(777.777"))
}

func IsSumOfCubes(s string) string {
	strSlice := strings.Split(s, "")
	var newNum string
	var intSlice []int
	for i := 0; i < len(strSlice); i++ {
		num, err := strconv.Atoi(strSlice[i])
		if err != nil {
			continue
		} else {
			if i == len(strSlice)-1 {
				newNum = newNum + fmt.Sprintf("%d", num)
				newInt, _ := strconv.Atoi(newNum)
				intSlice = append(intSlice, newInt)
				break
			} else {
				_, err := strconv.Atoi(strSlice[i+1])
				if err != nil {
					newNum = newNum + fmt.Sprintf("%d", num)
					newInt, _ := strconv.Atoi(newNum)
					intSlice = append(intSlice, newInt)
					newNum = ""
				} else if len(newNum) == 2 {
					newNum = newNum + fmt.Sprintf("%d", num)
					newInt, _ := strconv.Atoi(newNum)
					intSlice = append(intSlice, newInt)
					newNum = ""
				} else {
					newNum = newNum + fmt.Sprintf("%d", num)
				}
			}
		}
	}

	var result string
	var combinedTotal int
	for _, sepNum := range intSlice {
		var total int
		strInt := fmt.Sprintf("%d", sepNum)
		for _, digit := range strInt {
			digitInt, _ := strconv.Atoi(string(digit))
			add := digitInt * digitInt * digitInt
			total = total + add
		}
		if total == sepNum {
			combinedTotal = combinedTotal + sepNum
			result = result + fmt.Sprintf("%d ", sepNum)
		}
	}

	if result == "" {
		return "Unlucky"
	}

	return fmt.Sprintf("%s%d Lucky", result, combinedTotal)
}
