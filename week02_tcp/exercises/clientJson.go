package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

//make the shitty json one
//make one for large items
//large items + filename?????

type FileInfo struct {
	fileName string
	fileData []byte
}

var BUFFER_SIZE = 2048

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("Usage: sendFile filename01 filename02 ...")
	}

	currPath, _ := os.Getwd()

	files := make([]FileInfo, len(args))

	for index, file := range args {
		saveFileToMemory(&files[index], &currPath, &file)
	}

	addr := "127.0.0.1:8080"
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()

	if err != nil {
		log.Fatal("Error connecting:", err.Error())
	}

	fmt.Println("Connecting to TCP server in", addr)

	fmt.Println("Closing connection...")
}

func saveFileToMemory(currFile *FileInfo, currPath *string, fileName *string) {
	filePath := *currPath + "/" + *fileName

	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Println("File not found in current directory. It will not be sent:", *fileName)

		currFile.fileName = ""
		currFile.fileData = []byte("")

		return
	}

	stat, err := file.Stat()

	if err != nil {
		log.Println("File could not be read into memory. It will not be sent:", *fileName)

		currFile.fileName = ""
		currFile.fileData = []byte("")

		return
	}

	currFile.fileData = make([]byte, stat.Size())
	file.Read(currFile.fileData)
}
