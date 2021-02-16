package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)             // channel for sending strings
	done := make(chan bool)              // channel to OBSERVE closing of channel
	until := time.After(5 * time.Second) // channel to send timeout for app termination

	go send(msg, done) // start goroutine with channels for strings and closing message

	for { // endless loop
		select { // routes messages
		case m := <-msg: // assign received strings
			fmt.Println(m) // print received string
		case <-until: // wait for timeout
			done <- true                       // send bang to observer, func send() can close channel now
			time.Sleep(500 * time.Millisecond) // wait for send() termination
			return                             // func main() termination
		}
	}
}

func send(ch chan string, done <-chan bool) {
	for { // endless loop
		select {
		case <-done: // receives bang from observer
			println("Done")
			close(ch) // close channel
			return    // func send() termination
		default:
			ch <- "hello" // sends strings to channel
			time.Sleep(500 * time.Millisecond)
		}
	}
}
