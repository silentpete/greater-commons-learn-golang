// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."

// Range over the records, then range over the data in each record.

package main

import (
	"fmt"
)

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// use the "_" to throw out the "key"
	for _, x_value := range x {
		for _, value := range x_value {
			fmt.Println(value)
		}
	}

	// Todd made two seperate slices first
	// then added them together to make a slice of slice
	// then he used the key and value with print formatting
}
