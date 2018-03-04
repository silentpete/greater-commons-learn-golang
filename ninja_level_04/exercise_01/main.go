// Using a COMPOSITE LITERAL:
// create an ARRAY which holds 5 VALUES of TYPE int
// assign VALUES to each index position.
// Range over the array and print the values out.
// Using format printing
// print out the TYPE of the array

package main

import (
	"fmt"
)

func main() {
	fmt.Println("My Try at the Exercise")
	someAry := []int{1, 2, 3, 4, 5}
	fmt.Println(someAry)

	for x := range someAry {
		fmt.Println("index:", x, "value is:", someAry[x])
	}

	fmt.Printf("%T\n", someAry)

	fmt.Println("")

	fmt.Println("Todd's Example")

	x := [5]int{12, 23, 34, 45, 56}
	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)

}
