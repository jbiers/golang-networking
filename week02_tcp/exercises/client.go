package main

//JUST AGREE ON A PROTOCOL!!!!!!!!!!!!!!!!!!!!
import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var BUFFER_SIZE = 1024

func sendFileToServer(fileName *string, currPath *string, conn net.Conn) {
	filePath := *currPath + "/" + *fileName

	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Println("File not found in current directory:", *fileName)
		return
	}
	fileNameBuffer := make([]byte, 4)
	nextFileBuffer := make([]byte, 1)

	fileNameBuffer = []byte(*fileName)
	_, err = conn.Write(fileNameBuffer)
	if err != nil {
		fmt.Println("Failed to send filename to server:", *fileName)
		return
	}

	buffer := make([]byte, BUFFER_SIZE)

	time.Sleep(1 * time.Second)

	for {
		size, err := file.Read(buffer)

		if err == io.EOF {
			fmt.Println("Finished file transfer for:", *fileName)
			break
		}

		if err != nil {
			log.Println("Error reading data from file:", *fileName)
			return
		}

		_, err = conn.Write(buffer[:size])

		if err != nil {
			log.Fatal("Error transfering data from file:", *fileName)
			return
		}
	}
	nextFileBuffer = []byte("0")

	time.Sleep(1 * time.Second)

	_, err = conn.Write(nextFileBuffer)

}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("Usage: sendFile filename01 filename02 ...")
	}

	currPath, _ := os.Getwd()

	addr := "127.0.0.1:8080"
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()

	if err != nil {
		log.Fatal("Error connecting:", err.Error())
	}

	fmt.Println("Connecting to TCP server in", addr)

	for _, file := range args {
		sendFileToServer(&file, &currPath, conn)
	}

	fmt.Println("Closing connection...")
}
