package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() { // nameless F for catch the predicted panic and recover
		if err := recover(); err != nil {
			fmt.Printf("Catch the panic: %s (%T)\n", err, err)
		}
	}()
	fmt.Println("Heh")
	oopsy()
}

func oopsy() {
	panic(errors.New("bang the panic"))
}
