package httpserver

import (
	"fmt"
	"net"
	"os"
)

func Listen() (net.Listener, error) {
	listener, err := net.Listen("tcp", IPAddress() + ":" + Port()) // Listen for incoming connection
	if err != nil {
		return nil, err
	}
	fmt.Println("Listening on " + IPAddress() + ":" + Port())

	return listener, nil
}

func Read(conn net.Conn) (string, error) {
	buffer := make([]byte, 1024)
	readLength, error := conn.Read(buffer)
	return string(buffer[:readLength]), error
}

func Port() string {
	if os.Args > 2 {
		return os.Args[2]
	}
	return os.Args[1]
}

func IPAddress() string {
	if os.Args > 2 {
		return os.Args[1]	
	}
	return "0.0.0.0"
}
