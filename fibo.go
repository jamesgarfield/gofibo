package main

import "fmt"

func main() {
	n := GetFib(10)
	fmt.Printf("\nAnswer: %d\n", n)
}

func GetFib(n int64) int64 {
	c := make(chan int64)
	go fib(n, c)
	ans := <-c
	return ans
}

func fib(n int64, c chan int64) {
	switch {
	case n == 0:
		c <- 0
	case n <= 2:
		c <- 1
	default:
		cfib := make(chan int64)
		go fib(n-1, cfib)
		go fib(n-2, cfib)
		c <- (<-cfib + <-cfib)
	}
}
