// Create a program that shows the “if statement” in action.
// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(20)
	if x <= 5 {
		fmt.Println(x, "is less than 6")
	} else if x < 20 {
		fmt.Println(x, "is less than 20")
	} else {
		fmt.Println(x, "is not less than 6 and not less than 20")
	}
}
