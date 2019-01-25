package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

func getKeys(hosts map[string]IPAddr) []string {
	keys := make([]string, len(hosts))
	i := 0
	for k := range hosts {
		keys[i] = k
		i++
	}
	return keys
}

// TODO: Add a "String() string" method to IPAddr.
func (p IPAddr) String() string {
	//fmt.Println(int(p[0]))
	//	fmt.Println(len(p))
	res := ""
	for i := 0; i < len(p); i++ {
		res += strconv.Itoa(int(p[i]))
		if i > len(p)-1 {
			res += "."
		}
	}
	//fmt.Println(res)
	return res
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	//fmt.Println(hosts["loopback"])
	//fmt.Println(getKeys(hosts))

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
