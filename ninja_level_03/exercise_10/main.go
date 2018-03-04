// Write down what these print:
//  fmt.Println(true && true)
//  fmt.Println(true && false)
//  fmt.Println(true || true)
//  fmt.Println(true || false)
//  fmt.Println(!true)

package main

import (
	"fmt"
)

func main() {
	fmt.Println(true && true)  // should be true
	fmt.Println(true && false) // should be false
	fmt.Println(true || true)  // should be true
	fmt.Println(true || false) // should be true
	fmt.Println(!true)         // should be false
}
