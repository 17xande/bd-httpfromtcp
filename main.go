package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunkSize := 8
	buffer := make([]byte, chunkSize)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Printf("read: %s\n", string(buffer[:n]))
	}

}
