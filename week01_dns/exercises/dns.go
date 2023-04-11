package main

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
)

func IsUrlValid(u string) bool {

	if !strings.Contains(u, "http://") && !strings.Contains(u, "https://") {
		u = "https://" + u
	}

	_, err := url.ParseRequestURI(u)

	if err != nil {
		return false
	}

	return true
}

func PrintIps(u string) {
	fmt.Println("Searching IP addresses for:", u)

	if !IsUrlValid(u) {
		fmt.Println("URL is not valid.")
		return
	}

	ips, err := net.LookupIP(u)

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
