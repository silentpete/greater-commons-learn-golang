// Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

package main

import (
	"fmt"
)

func main() {
	//Previous Code
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println("original map:", m)

	// New for exercise 9
	// Composite Literal Example: Add to map
	// m[""] = []string{}
	m["gallerani_peter"] = []string{"learing", "rock climbing", "food"}

	// for mKeys1, mValues1 := range m {
	// 	fmt.Println("the key is:", mKeys1, "and values:", mValues1)
	// }

	for key1, value1 := range m {
		fmt.Println(key1)
		for key2, value2 := range value1 {
			fmt.Println(key2, "=", value2)
		}
	}
}
