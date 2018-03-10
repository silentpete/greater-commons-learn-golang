package main

import (
	"fmt"
)

func main() {
	for x := 33; x <= 122; x++ {
		fmt.Printf("Decimal: %d\n", x)
		fmt.Printf("- Binary: %b\n", x)
		fmt.Printf("- Hexadecimal: %#x\n", x)
		fmt.Printf("- Unicode: %#U\n", x)
	}
}
