package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	t1 := "12.255.56.1"
	// t2 := "123.456.789.0"
	fmt.Println(Is_valid_ip(t1))
}

func Is_valid_ip(ip string) bool {
	splitIp := strings.Split(ip, ".")
	if len(splitIp) != 4 {
		return false
	}
	for _, octet := range splitIp {
		if strings.Split(octet, "")[0] == "0" && len(octet) > 0 {
			return false
		}
		num, err := strconv.Atoi(octet)
		if err != nil || num > 255 || num < 0 {
			return false
		}
	}
	return true
}
