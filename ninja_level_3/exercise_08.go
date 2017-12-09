// Create a program that uses a switch statement with no switch expression specified.

package main

import (
	"fmt"
)

func main() {
	switch {
	case 1 == 1:
		fmt.Println("case true")
	case 1 == 2:
		fmt.Println("case false")
	}
}
