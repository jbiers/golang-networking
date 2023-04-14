package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// https://mrwaggel.be/post/golang-transfer-a-file-over-a-tcp-socket

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
	fileNameBuffer := make([]byte, 4)
	buffer := make([]byte, 1024)

	size, err := conn.Read(fileNameBuffer)
	fmt.Println(size)
	fmt.Println(string(fileNameBuffer[:size]))
	fmt.Printf("this is the filenamebuffer\n\n\n")
	if err != nil {
		log.Println("Failed to read filename sent by client:", err)
		return
	}

	fileName := string(fileNameBuffer[:size])

	size, err = conn.Read(buffer)

	fmt.Println(string(buffer[:size]))
	fmt.Println(fileName)

	file, err := os.Create(fileName)
	/*if err != nil {
		fmt.Println("Failed to create file:", fileName)
		return
	}*/
	defer file.Close()

	for {
		size, err := conn.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println("Error handling connection from:", conn.RemoteAddr())
			break
		}

		err = os.WriteFile(fileName, bytes.Trim(buffer[:size], "\x00"), 0644)
		//dont do this
		if err != nil {
			log.Println("Error saving file", fileName)
			break
		}
	}
	nextFileBuffer := make([]byte, 1)
	size, err = conn.Read(nextFileBuffer)

	fmt.Println(nextFileBuffer)

}
