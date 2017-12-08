// Create a program that shows the “if statement” in action.

package main

import (
	"fmt"
)

func main() {
	x := 8
	if x <= 5 {
		fmt.Println("less than 6")
	} else if x < 20 {
		fmt.Println("less than 20")
	} else {
		fmt.Println("not less than 6 and not less than 20")
	}
}
