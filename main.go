package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const port = ":42069"

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("could not open TCP %s: %s\n", port, err)
	}
	defer l.Close()
	fmt.Printf("Reading data from TCP %s\n", port)
	fmt.Println("=====================================")

	for {
		// Wait for a connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		log.Println("connection accepted")
		linesChan := getLinesChannel(conn)
		for line := range linesChan {
			fmt.Println("read:", line)
		}

		log.Println("connection closed")
	}
}

func getLinesChannel(c net.Conn) <-chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)
		currentLineContents := ""
		for {

			b := make([]byte, 8, 8)
			n, err := c.Read(b)
			if err != nil {
				if currentLineContents != "" {
					lines <- currentLineContents
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			str := string(b[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				lines <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
				currentLineContents = ""
			}
			currentLineContents += parts[len(parts)-1]
		}
	}()
	return lines
}
