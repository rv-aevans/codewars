package main

import (
	"fmt"
	"strconv"
	"strings"
)

var n = int64(12345)

func main() {
	fmt.Println(Smallest(n))
}

func Smallest(n int64) []int64 {
	numStr := strconv.FormatInt(n, 10)
	var ans int64 = n
	var ansI int64 = 0
	var ansJ int64 = 0
	for i := 0; i < len(numStr); i++ {
		for j := 0; j < len(numStr); j++ {
			arr := strings.Split(numStr, "")
			ch := arr[i]
			arr[i] = ""
			if i > j {
				arr[j] = ch + arr[j]
			} else {
				arr[j] = arr[j] + ch
			}
			newStr := strings.Join(arr, "")
			newInt, _ := strconv.ParseInt(newStr, 10, 64)
			if newInt < ans {
				ans = newInt
				ansI = int64(i)
				ansJ = int64(j)
			}
		}
	}
	return []int64{ans, ansI, ansJ}
}
