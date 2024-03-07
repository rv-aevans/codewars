package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	fmt.Println(IpsBetween("10.0.0.0", "10.0.0.50"))
}

func IpsBetween(start, end string) int {
	return int(binary.BigEndian.Uint32(net.ParseIP(end).To4()) - binary.BigEndian.Uint32(net.ParseIP(start).To4()))
}
