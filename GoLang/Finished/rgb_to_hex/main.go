package main

import "fmt"

func main() {
	res := convert(-20) + convert(275) + convert(125)
	fmt.Println(res)
}

func convert(n int) string {
	if n <= 0 {
		return "00"
	} else if n >= 255 {
		return "FF"
	} else {
		return fmt.Sprintf("%X", n)
	}
}
