package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	// args := os.Args[1:]                             // args taken from 1th, 0th is programm name
	// if result, err := Concat(args...); err != nil { // IF prepare condition; check condition {do the shit}
	// 	fmt.Printf("Error: %s\n", err)
	// } else {
	// 	fmt.Printf("Concatinated line: '%s...'\n", result) // \n on the end is to elliminate %-sign in output?
	// }

	//sweet rewrite. As Concat returns usefull value with error - empty string, no need to check and print something else
	rresult, _ := Concat(getArgs()...)
	fmt.Printf("Conrtgrtgnred strings: '%s'\n", rresult)
}

func getArgs() (parts []string) {
	return os.Args[1:]
}

// Concat the shit
// F gets any quantity of strings and returns one + error
func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings to work with") // empty string and new error
	}
	return strings.Join(parts, " "), nil // joined string and nil as error
}

func demoConcat() {
	s, err := Concat("Fuck", "me", "crab")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
