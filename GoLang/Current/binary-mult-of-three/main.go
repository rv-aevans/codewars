package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type finalNum int

func main() {
	n := finalNum(9)
	s := n.plusOne().divideByTwo().convertToString()

	t := reflect.TypeOf(s)
	fmt.Println(t)

	fmt.Println(s)
}

func (f *finalNum) plusOne() *finalNum {
	*f = *f + 1
	return f
}

func (f *finalNum) divideByTwo() *finalNum {
	*f = *f / 2
	return f
}

func (f *finalNum) convertToString() string {
	return strconv.Itoa(int(*f))
}
