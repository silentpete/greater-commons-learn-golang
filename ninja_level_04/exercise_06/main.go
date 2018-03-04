// Create a slice to store the names of all of the states in the United States of America. What is the length of your slice? What is the capacity? Print out all of the values, along with their index position in the slice, without using the range clause.

package main

import (
	"fmt"
)

func main() {
	// After listening to the instructions, Todd wants us to use "make"
	x := make([]string, 50, 50)
	fmt.Println(x)
	fmt.Println("length:", len(x))
	fmt.Println("capacity:", cap(x))

	x = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
		` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`,
		` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`,
		` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`,
		` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
		` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`,
		` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`,
		` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`,
		` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	// the below will cause more processing that the above
	// x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
	// 	` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`,
	// 	` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`,
	// 	` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`,
	// 	` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
	// 	` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`,
	// 	` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`,
	// 	` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`,
	// 	` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	// my original method of creating x
	// x := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`,
	//	` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`,
	//	` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`,
	//	` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`,
	//	` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
	//	` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`,
	//	` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`,
	//	` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`,
	//	` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,}
	fmt.Println(x)
	fmt.Println("length:", len(x))
	fmt.Println("capacity:", cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
