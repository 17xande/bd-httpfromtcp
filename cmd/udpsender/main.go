package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr := "localhost:42069"
	a, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, a)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Printf("Sending to %s. Type your message and press Enter to send. Press Ctrl+C to exit.\n", serverAddr)
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")
		s, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		i, err := conn.Write([]byte(s))
		if err != nil {
			panic(err)
		}
		fmt.Printf("sent %d bytes\n", i)
	}
}
