package main

import "fmt"

func main() {
	fmt.Println(Range(1, 4, 0))
}

func Range(s, e, st int) (res []int) {
	if e < s || s == e {
		return
	} else if st == 0 {
		for i := s; i < e-st; i++ {
			res = append(res, s)
		}
	} else {
		res = append(res, s)
		for s < e-st {
			s += st
			res = append(res, s)
		}
	}
	return
}
