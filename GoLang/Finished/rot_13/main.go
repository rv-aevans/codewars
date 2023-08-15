package main

import "fmt"

var (
	in  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	out = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
)

func main() {
	fmt.Println(Rot13("EBG13 rknzcyr."))
}

func Rot13(s string) string {
	rotMap := make(map[string]string)
	for i := 0; i < len(in); i++ {
		rotMap[string(in[i])] = string(out[i])
	}

	var res string

	for _, l := range s {
		v, ok := rotMap[string(l)]
		if ok {
			res += v
		} else {
			res += string(l)
		}
	}

	return res
}
