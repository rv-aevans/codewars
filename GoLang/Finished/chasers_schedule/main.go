package main

import (
	"fmt"
)

func main() {
	fmt.Println(Chaser(22, 10))
}

func Chaser(speed, time int) int {
	curMax := speed*time + speed
	for i := 1; i*2 < time; i++ {
		fmt.Println(curMax)
		fmt.Println(i)
		if curMax+(speed-i*3) > curMax {
			curMax = curMax + (speed - (i * 3))
		} else if curMax > speed*time+speed {
			return curMax
		}
	}
	return curMax
}
