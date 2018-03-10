// Create a value and assign it to a variable.
// Print the address of that value.
package main

import (
	"fmt"
)

func main() {
	variable := "value"
	fmt.Println("variable value is:\t", variable)
	fmt.Println("pointer address is:\t", &variable)
}
