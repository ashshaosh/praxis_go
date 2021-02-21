package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// ErrTimeout - returnable error on package level
var ErrTimeout = errors.New("Request timeout")

// ErrRejected - returnable error on package level
var ErrRejected = errors.New("Request rejected")
var random = rand.New(rand.NewSource(666))

// ParseError - Custom error type
type ParseError struct {
	Message    string //
	Line, Char int    // place in code
}

func (p *ParseError) Error() string {
	format := "%s on line: %d, char: %d"
	return fmt.Sprintf(format, p.Message, p.Line, p.Char)
}

func main() {
	response, err := SendRequest("Hi, crab")
	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying")
		response, err = SendRequest("Hi again, crab")
	} // this will run untill we get something else

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}

}

// SendRequest - response imitation
func SendRequest(req string) (string, error) {
	switch random.Int() % 333 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}
