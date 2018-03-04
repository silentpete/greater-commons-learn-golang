package main

import (
	"fmt"
)

type bacon int

var x bacon

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
