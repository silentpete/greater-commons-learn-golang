// Create and use an anonymous struct.

package main

import (
	"fmt"
)

func main() {
	a := struct {
		age    int
		height float64
		weight float64
	}{
		age:    37,
		height: 68.5,
		weight: 144.4,
	}
	fmt.Println("anonymous struct", a)
	fmt.Println(a.age)
	fmt.Println(a.height)
	fmt.Println(a.weight)

	// Challenge
	// using a map
	b := struct {
		m map[string]string
	}{
		m: map[string]string{
			"favMovie": "Matrix",
		},
	}
	fmt.Println("anonymous struct with map", b)

	// using a slice
	c := struct {
		age []int
	}{
		age: []int{
			37,
		},
	}
	fmt.Println("anonymous struct with slice", c)
}
