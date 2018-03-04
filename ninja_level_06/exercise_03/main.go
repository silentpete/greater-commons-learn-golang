// use defer example
package main

import (
	"fmt"
)

func main() {
	defer a()
	b()
}

func a() {
	fmt.Println("func a")
}

func b() {
	fmt.Println("func b")
}
