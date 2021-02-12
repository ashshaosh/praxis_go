package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go func() {
		fmt.Println("Inside goroutine")
		go echo(os.Stdin, os.Stdout)
	}()
	//go echo(os.Stdin, os.Stdout)
	time.Sleep(30 * time.Second)
	fmt.Println("Time is out")

	os.Exit(0)
}

func echo(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}
