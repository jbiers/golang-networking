package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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

	reader := bufio.NewReader(conn)

	fileName, err := reader.ReadBytes('\n')

	if err != nil {
		log.Println("error:", err)
	}

	fmt.Println(string(fileName))

	_, err = os.Create("KK")

	buffer := make([]byte, 1024)

	for {
		size, err := reader.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println("Error handling connection from:", conn.RemoteAddr())
			break
		}

		fmt.Println(buffer[:size])

		err = os.WriteFile("KK", buffer[:size], 0644)

		if err != nil {
			log.Println("Error saving file", fileName)
			break
		}
	}
}
