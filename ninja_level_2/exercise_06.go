package main

import (
	"fmt"
)

// example 1
const (
	// use iota print out last 4 years
	a = (2014 + iota)
	b
	c
	d
)

// example 2
const (
	e = (2020 + iota)
	f = (2020 + iota)
	g = (2020 + iota)
	h = (2020 + iota)

// so iota in the constant is an uilt in incrementor
)

func main() {
	fmt.Println("Example 1")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("Example 2")
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
