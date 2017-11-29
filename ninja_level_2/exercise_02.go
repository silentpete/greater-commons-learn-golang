package main

import (
	"fmt"
)

func main() {
	x := 2
	y := 4
	z := 4

	fmt.Println(z == y)
	fmt.Println(x <= y)
	fmt.Println(z >= x)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(y > x)

	// Woops, not exactly what was stated

	a := (10 == 20)
	b := (10 <= 20)
	c := (10 >= 20)
	d := (10 != 20)
	e := (10 < 20)
	f := (10 > 20)

	fmt.Println(a, b, c, d, e, f)
}
