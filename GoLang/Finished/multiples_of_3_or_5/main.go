package main

import "fmt"

func main() {
	fmt.Println(Multiple3And5(10))
}

func Multiple3And5(number int) (res int) {
	for i := 1; i < number; i++ {
		if i%5 == 0 || i%3 == 0 {
			res += i
		}
	}
	return
}
