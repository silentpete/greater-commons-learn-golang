// in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	num := 2
	wg.Add(num)
	for i := 0; i < num; i++ {
		go foo(i)
		// runtime.Gosched()
		wg.Done()
	}
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("exiting")
}

func foo(i int) {
	fmt.Println("hello from foo", i)
}
