package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	numOfGoroutines := 10
	numsAdded := 10

	// launch goroutine
	for i := 0; i < numOfGoroutines; i++ {
		fmt.Println("launching goroutine", i+1)
		go func(c chan int) {
			for n := 0; n < numsAdded; n++ {
				c <- n
			}
		}(c)
	}

	for total := 0; total < (numOfGoroutines * numsAdded); total++ {
		fmt.Println(total, <-c)
	}

}
