package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	// fmt.Println(checkPrime(3))
	fmt.Println(PrimeMaxLengthChain(176882))
	fmt.Println(time.Since(start))
}

func PrimeMaxLengthChain(n int) []int {
	orig := n

	if n%2 == 0 {
		n--
	}

	var primeList []int

	for i := n; i > 1; i -= 2 {
		if checkPrime(i) {
			primeList = append([]int{i}, primeList...)
		}
	}
	primeList = append([]int{2}, primeList...)

	possible := make(map[int]int)

	for j := 0; j < len(primeList); j++ {
		for k := len(primeList) / 2; k > j; k-- {
			sum := 0
			count := 0
			for _, val := range primeList[j:k] {
				sum += val
				count++
				if sum > n {
					break
				}
			}
			if sum > n {
				continue
			}
			if checkPrime(sum) && count > possible[sum] {
				possible[sum] = count
			}
		}
	}

	max := 0

	var res []int

	for k, v := range possible {
		if k == orig {
			continue
		}
		if possible[k] > max {
			max = v
			res = []int{k}
		} else if possible[k] == max {
			res = append(res, k)
		}
	}

	sort.Ints(res)

	return res
}

func checkPrime(n int) bool {
	if n < 2 {
		return false
	}

	sqRt := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqRt; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
