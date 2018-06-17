package main

import (
	"fmt"
)

func main() {
	quit := make(chan int)
	c := gen(quit)

	receive(c, quit)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive(c, quit <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-quit:
			fmt.Println("quiting")
			return
		}
	}
}
