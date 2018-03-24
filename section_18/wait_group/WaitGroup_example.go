package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// get OS
	fmt.Println("OS:\t", runtime.GOOS)
	// get Architecture
	fmt.Println("ARCH:\t", runtime.GOARCH)
	// get CPU count
	fmt.Println("CPUs:\t", runtime.NumCPU())
	// get go routines count
	fmt.Println("Go Routines:\t", runtime.NumGoroutine())

	wg.Add(1)
	go one()
	two()

	// get OS
	fmt.Println("OS:\t", runtime.GOOS)
	// get Architecture
	fmt.Println("ARCH:\t", runtime.GOARCH)
	// get CPU count
	fmt.Println("CPUs:\t", runtime.NumCPU())
	// get go routines count
	fmt.Println("Go Routines:\t", runtime.NumGoroutine())

	wg.Wait()
}

func one() {
	for i := 0; i < 10; i++ {
		fmt.Println("one:", i)
	}
	wg.Done()
}

func two() {
	for i := 0; i < 10; i++ {
		fmt.Println("two:", i)
	}
}
