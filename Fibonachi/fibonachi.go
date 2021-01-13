package main

import "fmt"

func fibonacciWithChannel(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

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
	c := make(chan int, 10)
	go fibonacciWithChannel(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	f := fibonacciWithSlices()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}

	for i := 9; i > 0; i-- {
		fmt.Println(fibonacciWithRecursion(i))
	}
}
