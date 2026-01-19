//stringer is like toString method in other languages
// it is used to define how a type should be represented as a string
// by implementing the Stringer interface from the fmt package.

package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	s := ""
	for idx, i := range ip {
		s += fmt.Sprintf("%d", i)
		if idx != len(ip)-1 {
			s += "."
		}
	}
	return s
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
