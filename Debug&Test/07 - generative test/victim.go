package gentest

import (
	"log"
	"strings"
)

func main() {
	Pad("lol", 6)
}

// Pad ...
func Pad(s string, max uint) string {
	log.Printf("Test lenght: %d, Str: %s\n", max, s)
	ln := uint(len(s))
	if ln > max {
		return s[:max-1] // must be s[:max] to pass test
	}
	s += strings.Repeat(" ", int(max-ln))
	return s
}
