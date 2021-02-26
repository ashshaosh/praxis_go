package main

import (
	"bytes"
	"fmt"
	"strings"
)

// How to concat two strings?
var sA string = "STRING!"
var sB string = "STRONG?"

func main() {
	// just + them
	s1 := sA + sB
	fmt.Printf("Joined: %s", s1)

	// you can even post something in between
	s2 := sA + " and " + sB
	fmt.Printf("Joined: %s", s2)

	// with +=
	s4 := sA
	s4 += sB
	fmt.Println(s4)

	// use buffer with bytes
	var b bytes.Buffer
	b.WriteString(sA)
	b.WriteString(sB)
	fmt.Println(b.String())

	// with fmt
	s3 := fmt.Sprintf("%s %s", sA, sB)
	fmt.Println(s3)

	// with Join F
	stringus := []string{sA, sB}
	s5 := strings.Join(stringus, " | ")
	fmt.Println(s5)

}
