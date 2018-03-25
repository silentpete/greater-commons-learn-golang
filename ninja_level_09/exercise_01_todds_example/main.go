package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	num := 2

	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("GRs:\t", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(num)

	go func() {
		fmt.Println("thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("thing two")
		wg.Done()
	}()

	fmt.Println("GRs:\t", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("exiting")
	fmt.Println("GRs:\t", runtime.NumGoroutine())
}
