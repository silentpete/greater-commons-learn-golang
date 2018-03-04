// Take the code from the previous exercise, then store the values of type
// person in a map with the key of last name. Access each value in the map.
// Print out the values, ranging over the slice.

package main

import (
	"fmt"
)

type person struct {
	first            string
	last             string
	favoriteIceCream []string
}

func main() {
	person1 := person{
		first: `Peter`,
		last:  `Gallerani`,
		favoriteIceCream: []string{
			"mint chocolate chip",
			"chocolate chocolate chip",
		},
	}
	person2 := person{
		first: `Nicole`,
		last:  `Smith`,
		favoriteIceCream: []string{
			`mint chocolate chip`,
			`birthday cake`,
			`chocolate peanut butter`,
		},
	}

	persons := map[string]person{
		person1.last: person1,
		person2.last: person2,
	}
	for _, v := range persons {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for _, l := range v.favoriteIceCream {
			fmt.Printf("\t%v\n", l)
		}
	}
}
