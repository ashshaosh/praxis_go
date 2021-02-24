package main

import (
	"fmt"
	"runtime"
)

func main() {
	foo()
}

func foo() {
	bar()
}

func bar() {
	buf := make([]byte, 1024) // make bufer
	runtime.Stack(buf, false) // write stack into bufer. False to trace all running goroutines
	fmt.Printf("Trace:\n %s\n", buf)
	//debug.PrintStack()
}
