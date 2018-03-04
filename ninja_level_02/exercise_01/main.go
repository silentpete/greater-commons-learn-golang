package main

import (
	"fmt"
)

var x int = 42

func main() {
	fmt.Printf("Decimal: %d\n", x)
	fmt.Printf("Binary: %b\n", x)
	fmt.Printf("Hexadecimal: %#x\n", x)
}
