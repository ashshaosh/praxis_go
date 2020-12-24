// Newton's methods for calculating the square root in golang

package main

import (
	"fmt"
	"math"
)

// Sqrt ...
func Sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}

	z := float64(1)
	var prev float64

	for math.Abs(z-prev) > .000000001 {
		prev = z
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(32))
	fmt.Println(math.Sqrt(32))
}
