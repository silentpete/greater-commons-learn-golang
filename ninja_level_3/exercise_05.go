// Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.

package main

import (
	"fmt"
)

func main() {
	for x := 10; x <= 100; x++ {
		if x%4 == 0 {
			fmt.Println(x)
		}
	}
}
