package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:8080"
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Println("Error connecting:", err.Error())
		return
	}

	defer conn.Close()

	fmt.Println("Connecting to TCP server in", addr)
	message := "Test message"

	_, err = conn.Write([]byte(message))

	if err != nil {
		log.Println("Error sending message:", err.Error())
		return
	}
}
