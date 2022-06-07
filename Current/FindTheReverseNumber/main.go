package main

import "fmt"

func main() {
	fmt.Println(FindReverseNumber(100000000))
}

func FindReverseNumber(n uint64) uint64 {
	var resCount uint64 = 0
	for i := 0; i >= 0; i++ {
		if fmt.Sprintf("%d", i) == reverseString(i) {
			resCount++
		}
		if resCount == n {
			return uint64(i)
		}
		fmt.Println(resCount)
	}
	return 0
}

func reverseString(n int) string {
	runes := []rune(fmt.Sprintf("%d", n))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
