package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(NbMonths(18000, 32000, 1500, 1.25))
}

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	b, g, p := float64(startPriceOld), float64(startPriceNew), float64(1)
	for i := 0; ; i++ {
		if i == 1 {
			p -= (percentLossByMonth / 100)
		} else if i%2 == 0 && i != 0 {
			p -= .005
		}
		b, g = b*p, g*p
		if (int(b) + i*savingperMonth) >= int(g) {
			return [2]int{i, int(math.Round((b + float64(i*savingperMonth)) - g))}
		}
	}
}
