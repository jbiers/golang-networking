package main

import (
	"fmt"
	"net"
	"os"
)

func PrintIps(h string) {
	fmt.Println("Searching IP addresses for:", h)

	ips, err := net.LookupIP(h)

	if err != nil {
		fmt.Println(err)
		fmt.Println("")

		return
	}

	for _, ip := range ips {
		fmt.Println("IP address found:", ip)
	}

	fmt.Println("")
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		fmt.Println("Usage: go run dns.go hostname_01 hostname_01...")
	}

	for _, host := range argsWithoutProg {
		PrintIps(host)
	}
}
