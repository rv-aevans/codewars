package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(SpinningRings(16777216, 14348907))
}

func SpinningRings(innerMax, outerMax uint64) uint64 {
	s := time.Now()
	start := math.Abs(float64(outerMax) - float64(innerMax))
	i, o, c := innerMax, uint64(1), uint64(start)
	for i != o {
		c++
		i = (i + innerMax) % (innerMax + 1)
		o = (o + 1) % (outerMax + 1)
	}
	fmt.Println(time.Since(s))
	return c
}
