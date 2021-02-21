package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("Cant /0")

func main() {
	//panic("Heh, how are you?") // instant crash
	//panic(errors.New("Bad shit happened"))
	fmt.Println("Dvd 1 by 0")
	_, err := precheckDivide(1, 0) // proper handling
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println("Dvd 2 by 0")
	divide(2, 0) // panic and crash
}

func precheckDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return divide(a, b), nil
}

func divide(a, b int) int {
	return a / b
}
