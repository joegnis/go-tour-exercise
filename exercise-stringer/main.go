/*
Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
*/
package main

import (
	"strconv"
	"strings"
	"fmt"
)

type IPAddr [4]byte

func (addr IPAddr) String() string {
	string_addr := make([]string, 4)
	for i, byte_addr := range addr {
		string_addr[i] = strconv.Itoa(int( byte_addr ))
	}
	return strings.Join(string_addr, ".")
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
