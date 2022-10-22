package main

import "fmt"

func main() {
	fmt.Println(f("MCDLXX"))
}

func f(s string) (d int) {
	for i := 0; i < len(s); i++ {
		fmt.Println(d)
		if i == len(s)-1 && (dMap[string(s[i])]) <= (dMap[string(s[i-1])]) {
			d += dMap[string(s[i])]
			return
		} else if i == len(s)-1 {
			return
		}
		if dMap[string(s[i])] >= dMap[string(s[i+1])] {
			d += dMap[string(s[i])]
		} else {
			d += (dMap[string(s[i+1])] - dMap[string(s[i])])
			i += 1
		}
	}
	return
}

var dMap = map[string]int{
	"M": 1000,
	"D": 500,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}
