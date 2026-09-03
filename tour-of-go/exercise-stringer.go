package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var result strings.Builder
	for i, n := range ip {
		if i > 0 {
			result.WriteString(".")
		}
		result.WriteString(strconv.Itoa(int(n)))
	}
	return result.String()
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
