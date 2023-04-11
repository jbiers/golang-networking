package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("Usage: go run client.go filename01 filename02 ...")
	}

	currPath, _ := os.Getwd()

	file, err := os.Open(currPath + "/" + args[0])
	defer file.Close()

	if err != nil {
		log.Fatal("Could not open file:", err.Error())
	}

	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		log.Fatal("Could not read file:", err.Error())
	}

	fmt.Println(n)

	addr := "127.0.0.1:8080"
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()

	if err != nil {
		log.Fatal("Error connecting:", err.Error())
	}

	fmt.Println("Connecting to TCP server in", addr)

	_, err = conn.Write([]byte(buffer))

	if err != nil {
		log.Fatal("Error sending message:", err.Error())
	}
}
