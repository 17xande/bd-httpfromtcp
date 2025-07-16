package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const inputFilePath = "messages.txt"

func main() {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", inputFilePath, err)
	}
	defer file.Close()
	chunkSize := 8
	buffer := make([]byte, chunkSize)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatalf("error: %s\n", err.Error())
		}
		fmt.Printf("read: %s\n", string(buffer[:n]))
	}

}
