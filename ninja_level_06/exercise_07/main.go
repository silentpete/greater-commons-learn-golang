// Assign a func to a variable, then call that func
package main

import (
	"fmt"
)

func main() {
	x := sayHello
	x()
	fmt.Printf("type: %T\n", x)
}

func sayHello() {
	fmt.Println("hello")
}
