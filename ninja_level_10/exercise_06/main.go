package main

import (
	"fmt"
)

func main() {
	// bidirectional channel
	fmt.Println("make c")
	c := make(chan int)

	// put numbers on channel
	fmt.Println("add nums to c")
	addNums(c)

	// print all numbers on channel
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("gonna exit")
}

// addNums adds numbers to the channel
func addNums(c chan int) chan int {
	fmt.Println("send out goroutine adding to c")
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		fmt.Println("gonna close c")
		close(c)
	}()
	fmt.Println("gonna return c")
	return c
}
