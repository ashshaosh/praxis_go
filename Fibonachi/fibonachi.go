package main

import "fmt"

func fibonacciWithSlices() func(int) int {
	return func(x int) int {
		fnumbers := make([]int, x+1, x+2)
		if x < 2 {
			fnumbers = fnumbers[0:2]
		}
		fnumbers[0] = 0
		fnumbers[1] = 1

		for i := 2; i <= x; i++ {
			fnumbers[i] = fnumbers[i-1] + fnumbers[i-2]
		}
		return fnumbers[x]
	}
}

func fibonacciWithRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciWithRecursion(n-1) + fibonacciWithRecursion(n-2)
}

func main() {
	f := fibonacciWithSlices()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}

	for i := 9; i > 0; i-- {
		fmt.Println(fibonacciWithRecursion(i))
	}
}
