// Anonymous self-executing func
package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Hello World!")
	}()
	func(i int) {
		fmt.Println("passed int is:", i)
	}(8)
}
