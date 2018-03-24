package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CUPs:", runtime.NumCPU())
	fmt.Println("Go routines:", runtime.NumGoroutine())
	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("counter:", counter)
	fmt.Println("Go routines:", runtime.NumGoroutine())
}
