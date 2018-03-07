// Recursion
package main

import (
	"fmt"
)

func main() {
	fmt.Println(3 * 2 * 1)
	fmt.Println(factorial(5))
	fmt.Println(loopFactorial(5))
}

func factorial(i int) int {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}

func loopFactorial(i int) int {
	value := 1
	for n := i; n > 0; n-- {
		value = value * n
	}
	return value
}
