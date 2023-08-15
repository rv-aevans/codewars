package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func main() {
	// fmt.Println(GetMinBase(8065425431))
	fmt.Println(math.Sqrt(9295997013522923649))
	// fmt.Println(math.Sqrt(8065425431))
}

func GetMinBase(n uint64) uint64 {
	res := uint64(0)
	bn := big.NewInt(int64(n))
	for i := 2; i <= 58; i++ {
		if numDigMatch(bn.Text(i)) {
			return uint64(i)
		}
	}

	if res != 0 {
		return res
	}

	s := fmt.Sprintf("%d", n)
	if string(s[len(s)-1]) == "1" {
		ns := "1"
		for i := 2; i < len(s); i++ {
			if string(s[len(s)-i]) == "0" {
				ns = "0" + ns
			} else if len(ns) > 2 {
				ns = string(s[len(s)-i]) + ns
				break
			} else {
				return n - 1
			}
		}
		newN, _ := strconv.Atoi(ns)
		return uint64(newN - 1)
	} else {
		return n - 1
	}
}

func numDigMatch(s string) bool {
	for i, v := range s {
		if i == len(s)-1 {
			break
		}
		if string(v) != string(s[i+1]) {
			return false
		}
	}
	return true
}
