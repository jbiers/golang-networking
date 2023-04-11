package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:8080"

	listener, err := net.Listen("tcp", addr)
	defer listener.Close()

	if err != nil {
		log.Fatal("Error creating listener:", err.Error())
	}

	fmt.Println("TCP Server listening on", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err.Error())
			return
		}

		fmt.Println("New connection from", conn.RemoteAddr())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)

		if err != nil && err != io.EOF {
			log.Println("Error handling connection from:", conn.RemoteAddr())
			return
		}

		if err == io.EOF {
			return
		}

		fmt.Println(string(buffer[:n]))
	}
}
