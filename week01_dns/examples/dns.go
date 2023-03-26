package main

import (
	"fmt"
	"net"
)

func main() {
	ips, err := net.LookupIP("45d")

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, ip := range ips {
		fmt.Println("IP encontrado: ", ip)
	}
}
