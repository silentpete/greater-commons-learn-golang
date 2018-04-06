package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	gr := 100

	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			fmt.Println("c:", counter)
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter:", counter)
}
