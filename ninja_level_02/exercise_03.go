package main

import (
	"fmt"
)

// untyped constant
const a = 24
// typed constant
const b int = 4

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
