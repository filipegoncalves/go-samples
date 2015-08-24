
package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	res := fmt.Sprintf("%v", ip[0])
	for _, b := range ip[1:] {
		res += fmt.Sprintf(".%v", b)
	}
	return res
}

func main() {
	addrs := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
