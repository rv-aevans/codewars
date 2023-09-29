package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(RectangleRotation(6, 4))
}

type point struct {
	x float64
	y float64
}

func RectangleRotation(a, b int) int {
	c, s := math.Cos(.785), math.Sin(.785)
	x1, y1 := (float64(a/2)*c - float64(b/2)*s), (float64(a/2)*s + float64(b/2)*c)
	x2, y2 := (float64(a/2)*c - float64(-b/2)*s), (float64(a/2)*s + float64(-b/2)*c)
	x3, y3 := -x1, -y1
	fmt.Println(x1, y1)
	fmt.Println(x2, y2)
	fmt.Println(x3, y3)

	fmt.Println(math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)))
	fmt.Println(math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2)))

	return 0
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func getPointsOnLine(p, q point) int {
	if p.x == q.x {
		return int(math.Abs(p.y-q.y)) - 1
	} else if p.y == q.y {
		return int(math.Abs(p.x-q.x)) - 1
	} else {
		return gcd(int(math.Abs(p.x-q.x)), int(math.Abs(p.y-q.y))) - 1
	}
}
