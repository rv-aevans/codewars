package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println(FindAll(27, 17))
}

func FindAll(sumDig, digs int) []int {
	min, max, count, resSlice := "1", "9", 0, []int{}
	var wg sync.WaitGroup
	for i := 1; i < digs; i++ {
		min += "0"
		max += "9"
	}
	// intMin, _ := strconv.Atoi(min)
	intMax, _ := strconv.Atoi(max)
	wg.Add(intMax)
	for i := 0; i <= intMax; i++ {
		go func(i int) {
			defer wg.Done()
			total := 0
			split := strings.Split(fmt.Sprintf("%d", i), "")
			for j, str := range split {
				if j != 0 && split[j-1] > str {
					total = 0
					break
				}
				num, _ := strconv.Atoi(str)
				total += num
				if total > sumDig {
					break
				}
			}
			if total == sumDig {
				resSlice = append(resSlice, i)
				count++
			}
		}(i)
	}

	wg.Wait()

	if len(resSlice) == 0 {
		return nil
	}
	return []int{count, resSlice[0], resSlice[len(resSlice)-1]}
}
