package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var input []int

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i <= 4; i++ {
		var level []int
		n := math.Pow(2, float64(i))
		for j := int(n); j > 0; j-- {
			rn := r.Intn(int(n) + 1)
			level = append(level, rn)
			input = append(input, rn)
		}
		fmt.Println(level)
	}

	res := 1
	fmt.Println([]int{1})
	for i := 1; res < len(input); i++ {
		var level []int
		k := int(math.Pow(2, float64(i)))
		s := (k + res - 1)
		for j := 0; j < k; j++ {
			level = append(level, input[s-j])
		}
		res += k
		fmt.Println(level)
	}
}
