package main

import (
	"fmt"
)

func main() {

	x := 42

	fmt.Printf("%d\t", x)
	fmt.Printf("%b\t", x)
	fmt.Printf("%#x\n", x)


	// bit shift
	y := (x << 1)

	fmt.Printf("%d\t", y)
	fmt.Printf("%b\t", y)
	fmt.Printf("%#x\n", y)
}
