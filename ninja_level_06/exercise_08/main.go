// Create a func which returns a func
// assign the returned func to a variable
// call the returned func
package main

import (
	"fmt"
)

// var x func()

func main() {
	passingText := sFunc()
	fmt.Println(passingText)

	x := someFunc()

	stxt := x()
	fmt.Println(stxt)
}

func sFunc() string {
	return "Hello World!"
}

func someFunc() func() string {
	fmt.Println("inside some func")
	return func() string {
		return "from returned func"
	}
}

// I did not understand this fully, did half the work, then was the prior video about returning a func
