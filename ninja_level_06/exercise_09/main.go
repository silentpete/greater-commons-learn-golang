// A “callback” is when we pass a func into a func as an argument.
// For this exercise, pass a func into a func as an argument
package main

import (
	"fmt"
)

func main() {
	name := "silentpete"
	wLength := wordLength(name)
	fmt.Println(name, "has", wLength, "characters")

	wDiv := divByTwo(wordLength, name)
	fmt.Println("modulo 2?", wDiv)
}

func wordLength(s string) int {
	return len(s)
}

func divByTwo(si func(s1 string) int, s2 string) string {
	sNum := si(s2)
	if sNum%2 == 0 {
		return "true"
	}
	return "false"
}
