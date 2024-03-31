package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
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

		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		len, err := conn.Read(buf)

		if err == io.EOF {
			break
		}

		fullBuf := buf[:len]
		// %q is to safely escape the string
		message := fmt.Sprintf("%q", fullBuf)
		fmt.Println("received:", message)

		parseRequest(fullBuf)

		if err != nil {
			fmt.Println("failed to read buffer", err)

		}

		conn.Write([]byte("+PONG\r\n"))
	}
}

type Type byte

const (
	Integer = ':'
	String  = '+'
	Bulk    = '$'
	Array   = '*'
	Error   = '-'
)

func parseRequest(b []byte) int {
	if len(b) == 0 {
		return 0
	}
	messageType := Type(b[0])

	if messageType == Array {
		fmt.Printf("array message; %q", b[1:])
		parseRequest(b[:1])
	}

	return 0
}
