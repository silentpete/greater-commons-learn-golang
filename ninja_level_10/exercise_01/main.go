package main

import (
	"fmt"
)

func main() {
	// buffered channel
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
	a()
}

func a() {
	// unbuffered channel
	c := make(chan int)

	// because the channel blocks, the channel is sent out on it's own goroutine
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
