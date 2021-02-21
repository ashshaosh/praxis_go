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

// Concat the shit
// F gets any quantity of strings and returns one + error
func Concat(parts ...string) (string, error) {
	var e error = nil // nil error
	s := ""           // empty string
	if len(parts) == 0 {
		e = errors.New("No strings to work with") // !nil error
		fmt.Printf("well... %s\n", e)
		//return "", e // empty string and new error
	}
	s = strings.Join(parts, " ") // fill string with args
	return s, e
}

func getArgs() (parts []string) {
	return os.Args[1:]
}

func demoConcat() {
	s, err := Concat("Fuck", "me", "crab")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
