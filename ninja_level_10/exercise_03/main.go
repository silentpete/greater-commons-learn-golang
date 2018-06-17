package main

import (
	"fmt"
)

func main() {
	// call gen
	c := gen()
	// call receive and pass it parameter 'c'
	receive(c)

	fmt.Println("about to exit")
}

// gen will return a receive/out channel of Type int.
func gen() <-chan int {
	// make bidirectional channel
	c := make(chan int)
	// send out goroutine
	go func() {
		// put 0 - 9 on channel
		for i := 0; i < 10; i++ {
			c <- i
		}
		// close channel from blocking after finished loop
		close(c)
	}()

	return c
}

// receive takes a receive/out channel of Type int.
func receive(c <-chan int) {
	// range over channel 'c'
	for v := range c {
		// print out each value
		fmt.Println(v)
	}
}
