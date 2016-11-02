package httpserver

import (
	"fmt"
	"io/ioutil"
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
	result, error := ioutil.ReadAll(conn)
	return string(result), error
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
