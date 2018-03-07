// Closure is when we have “enclosed” the scope of a variable in some code block.
// For this hands-on exercise, create a func which “encloses” the scope of a variable
package main

import (
	"fmt"
)

var x int

func main() {
	fmt.Println(x) // x exists in all of main package
	x++
	fmt.Println(x)
	foo()
	d := incrementor()
	fmt.Println(d())
	fmt.Println(d())
}

func foo() {
	b := 8
	fmt.Println("b =", b, ", only is in foo")
}

func incrementor() func() int {
	// this memory space will keep 'c' from being used outside this func, but can be assigned outside and used
	c := 0
	return func() int {
		c++
		return c
	}
}

// used todds func
