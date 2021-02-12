package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdin, os.Stdout, "outside")
	go func() {
		fmt.Println("Inside goroutine")
		go echo(os.Stdin, os.Stdout, "inside")
	}()

	time.Sleep(30 * time.Second)
	fmt.Println("Time is out")

	os.Exit(0)
}

func echo(in io.Reader, out io.Writer, where string) {
	fmt.Println(where)
	io.Copy(out, in)
}
