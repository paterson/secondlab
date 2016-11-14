package httpserver

import (
	"fmt"
	"net"
	"os"
)

func Listen() (net.Listener, error) {
	listener, err := net.Listen("tcp", "0.0.0.0:"+Port()) // Listen for incoming connection
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
	return os.Args[1]
}

func IPAddress() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return string(ipv4)
		}
	}
	return "0.0.0.0"
}
