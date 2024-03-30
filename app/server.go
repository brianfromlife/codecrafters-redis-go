package main

import (
	"fmt"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {

		conn, err := l.Accept()

		if err != nil {
			fmt.Println("failed to accept")
			os.Exit(1)
		}

		buf := make([]byte, 1024)

		_, err = conn.Read(buf)

		if err != nil {
			fmt.Println("failed to read buffer", err)
			os.Exit(1)
		}

		conn.Write([]byte("+PONG\r\n"))
		conn.Close()
	}

}
