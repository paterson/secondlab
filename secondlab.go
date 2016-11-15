package main

import (
	"fmt"
	"github.com/paterson/secondlab/httpserver"
	"github.com/paterson/secondlab/workermanager"
	"net"
	"os"
	"strings"
)

const (
	NUMBER_OF_WORKERS = 3
	HELO_TEXT         = "HELO text"
	KILL_SERVICE      = "KILL_SERVICE"
)

func main() {
	// Listen for incoming connection
	listener, err := httpserver.Listen()
	checkError(err)

	jobs := make(chan workermanager.Job, 100)
	workermanager.Start(NUMBER_OF_WORKERS, jobs)

	defer listener.Close() // Close the listener when the application closes.
	for {
		connection, err := listener.Accept() // Accept incoming connection
		checkError(err)
		job := workermanager.Job{Action: handleRequest, Conn: connection}
		jobs <- job
	}
}

func handleRequest(connection net.Conn) {
	message, err := httpserver.Read(connection)
	checkError(err)

	fmt.Println("Request Received:", message)
	if message.HasPrefix(HELO_TEXT) {
		fmt.Println("Here")
		respondToHello(connection)
	} else if message == KILL_SERVICE {
		killService(connection)
	} else {
		doNothing(connection)
	}
}

func killService(connection net.Conn) {
	connection.Close()
	os.Exit(0)
}

func respondToHello(connection net.Conn) {
	response := fmt.Sprintf("HELO text\nIP:%s\nPort:%s\nStudentID:12305503\n", httpserver.IPAddress(), httpserver.Port())
	connection.Write([]byte(response))
	connection.Close()
}

func doNothing(connection net.Conn) {
	connection.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
