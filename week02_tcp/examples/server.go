package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:8080"

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Println("Error creating listener:", err.Error())
		return
	}

	defer listener.Close()

	fmt.Println("TCP Server listening on", addr)

	conn, err := listener.Accept()
	if err != nil {
		log.Println("Error accepting connection:", err.Error())
	}

	fmt.Println("New connection from", conn.RemoteAddr())

	err = handleConnection(conn)
	if err != nil {
		log.Println("Error accepting connection:", err.Error())
	}
}

func handleConnection(conn net.Conn) error {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	if err != nil {
		return err
	}

	fmt.Println("Received data:", string(buffer[:n]))

	return nil
}
