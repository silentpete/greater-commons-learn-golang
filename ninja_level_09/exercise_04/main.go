package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	gr := 100

	var wg sync.WaitGroup
	wg.Add(gr)

	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			m.Lock()
			fmt.Println("c:", counter)
			v := counter
			// runtime.Gosched()
			v++
			counter = v
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter:", counter)
}
