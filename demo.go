package main

import (
	"fmt"
)

func fib(x int, c chan int) {
	if x == 0 {
		c <- 0
	} else if x == 1 {
		c <- 1
	} else {
		c1 := make(chan int)
		go fib(x-1, c1)
		go fib(x-2, c1)
		val1, val2 := <-c1, <-c1
		c <- (val1 + val2)
	}
}

func main() {
	fmt.Println("Hello, GO.")

	recv := make(chan int)
	go fib(15, recv)
	val := <-recv

	fmt.Println(val)
}
