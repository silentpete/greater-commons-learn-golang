// Create a map with a key of TYPE string which is a person’s “last_first”
// name, and a value of TYPE []string which stores their favorite things. Store
// three records in your map. Print out all of the values, along with their
// index position in the slice.
//
//   `bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
//   `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
//   `no_dr`, `Being evil`, `Ice cream`, `Sunsets`

package main

import (
	"fmt"
)

func main() {
	//I couldn't figure it out, so I had to watch Todds video

	// create a map of "m" with short initiator ":="
	// If we want more than one value per key, it will need to be a slice of TYPE
	// map with key of TYPE string, and values of slice of TYPE string
	// Composite Literal Example: Create Map
	// m:= map[string][]string{}
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println(m)

	for key1, value1 := range m {
		fmt.Println(key1)
		for key2, value2 := range value1 {
			fmt.Println(key2, "=", value2)
		}
	}

}
