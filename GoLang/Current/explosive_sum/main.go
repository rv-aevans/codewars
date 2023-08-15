package main

import (
	"fmt"
	"math"
)

func main() {
	sqRt := math.Sqrt(31)
	// fmt.Println(2.8282 * 8)
	fmt.Println(sqRt)
	mod := int(math.Floor(sqRt)) % 31
	fmt.Println(mod)
	fmt.Println(math.Pow(float64(mod), sqRt))
}
