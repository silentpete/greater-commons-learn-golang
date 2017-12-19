// Create your own type “person” which will have an underlying type of “struct”
// so that it can store the following data:
//   first name
//   last name
//   favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the
// elements in the slice which stores the favorite flavors.

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

	fmt.Printf("%v\n%v\n", person1, person2)

	people := []person{
		person1,
		person2,
	}

	for _, personValue := range people {
		fmt.Println(personValue.first, "likes:")
		for _, person1value := range personValue.favoriteIceCream {
			fmt.Printf("\t%v\n", person1value)
		}
	}
}
