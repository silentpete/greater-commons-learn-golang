// a function syntax example
// func (receiver) identifier(parameters) returns { code }
package main

import (
	"fmt"
)

func main() {
	x := firstExample()
	fmt.Println("first example:", x)
	fmt.Println("first example passing in func:", firstExample())
	fmt.Println("second example:", secondExample(100))
	fmt.Println("second example passing in func:", secondExample(firstExample()))
}

// keyword        = func
// firstExample() = identifier
// int            = return
func firstExample() int {
	x := 8
	return x
}

// keyword    = func
// itentifier = n int
// the int passed in will be assigned to variable n
// return     = int
func secondExample(n int) int {
	x := n
	x++
	return x
}
