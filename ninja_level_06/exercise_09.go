// A “callback” is when we pass a func into a func as an argument. For this exercise,
// pass a func into a func as an argument
package main

import (
	"fmt"
)

func main() {
	word := "hello"
	fmt.Println(say(word))
	fmt.Println(sayPlus(say(word)))
}

func say(w string) string {
	return w
}
func sayPlus(w string) int {
	return len(w)
}

// not really as cool as todd's example and I'm not sure I'm doing callbacks correctly
