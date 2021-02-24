package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1902") // make connection to remote server
	if err != nil {
		panic("Failed to connect")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "CRAB ", f) // make logger with connection as writable object
	logger.Println("Usual shit")        // write log to remote server
	logger.Panicln("Panic on board")
}

// to test this use
// Netcat https://g.co/kgs/6wJRxj
// Netcat is a computer networking utility for reading from and writing to network connections using TCP or UDP. The command is designed to be a dependable back-end that can be used directly or easily driven by other programs and scripts.

// nc -lk 1902
// -l listen
// -k keep-open
// -p port on Debian version
