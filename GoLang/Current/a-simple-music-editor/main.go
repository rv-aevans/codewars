package main

import "fmt"

func main() {
	// fmt.Println(solution([]int{1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
	fmt.Println(solution([]int{1, 1, 1, 3}))
}

func solution(raw []int) string {
	for i := 0; i < len(raw); i++ {
		count, index := checkConsecutive(raw, raw[i], 1, i)
		fmt.Println("Count: ", count)
		fmt.Println("Index: ", index)
		i = index - 1
	}
	return "hey"
}

func checkConsecutive(raw []int, cur, count, index int) (int, int) {
	if cur == raw[index+1] {
		return checkConsecutive(raw, cur, count+1, index+1)
	}
	return count, index
}
