package main

import "fmt"

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	fmt.Println(OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123"))
// }

// func OrderWeight(strng string) string {
// 	splitStr := strings.Split(strng, " ")

// 	resMap := make(map[string]int)
// 	var weights []int

// 	for _, val := range splitStr {
// 		origStr, weight := getWeight(val)
// 		weights = append(weights, weight)
// 		resMap[origStr] = weight
// 	}

// 	sort.Ints(weights)

// 	fmt.Println(weights)
// 	fmt.Println(resMap)

// 	var resString string

// 	for _, weight := range weights {
// 		var count int
// 		var keys []string
// 		for k, v := range resMap {
// 			if weight == v {
// 				keys = append(keys, k)
// 				count++
// 			}
// 		}

// 		if count == 1 {
// 			resString = resString + keys[0] + " "
// 		} else {
// 			sort.Strings(keys)
// 			for _, key := range keys {
// 				resString = resString + key + " "
// 			}
// 		}
// 	}

// 	return strings.TrimSuffix(resString, " ")
// }

// func getWeight(str string) (string, int) {
// 	splitNum := strings.Split(str, "")

// 	var weight int

// 	for _, dig := range splitNum {
// 		intDig, _ := strconv.Atoi(dig)
// 		weight += intDig
// 	}

// 	return str, weight
// }

func main() {
	str := "what"

	for _, v := range str {
		fmt.Println(int(v))
	}
}
