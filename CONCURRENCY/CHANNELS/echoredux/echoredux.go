package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// pause with timer and channel demo:
	time.Sleep(1 * time.Second)          // create (useless) pause with Sleep
	sleep := time.After(1 * time.Second) // create (useless) pause with channel
	<-sleep                              // function will sleep untill message from channel SLEEP to no one, it's just some point to pause function
	// this will pause execution for 2sec

	done := time.After(30 * time.Second) // create one-way channel with time.Time object
	echo := make(chan []byte)            // create channel with MAKEf

	go readStdin(echo) // goroutine started with channel ECHO

	for { // endless loop for select to able wait for messages
		select { // routes signals from channels
		case buf := <-echo: // reads values from ECHO channel and assigns it to var
			os.Stdout.Write(buf)
		case <-done: // reacts to message in channel DONE and terminate's app
			fmt.Println("Time out")
			os.Exit(0)
		}
	} // w/o default select'll lock execution until one of the cases receives message

}

func readStdin(out chan<- []byte) { // receives channel to store data
	for {
		data := make([]byte, 1024)  // object to store data
		l, _ := os.Stdin.Read(data) // store input-data in object
		if l > 0 {                  // if bytes not zero
			out <- data // pass data-object to channel
		}
	}
}
