// Print every rune code point of the uppercase alphabet three times.
// Example
// 65
//	U+0041 'A'
//	U+0041 'A'
//	U+0041 'A'

package main

import (
	"fmt"
)

func main() {
	starting_rune := 65
	ending_rune := 90
	for i := starting_rune; i <= ending_rune; i++ {
		fmt.Println(i)
		for x := 1; x <= 3; x++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
