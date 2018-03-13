// marshal the []user to JSON
// Each exported struct field becomes a member of the object.
// fields must be exported identifiers. This means the field name start with an upper case
package main

import (
	"encoding/json"
	"fmt"
)

// create a structure
type person struct {
	First string
	Age   int
}

func main() {
	// create a person with the created structure
	p1 := person{
		First: "James",
		Age:   32,
	}

	p2 := person{
		First: "Moneypenny",
		Age:   27,
	}

	p3 := person{
		First: "M",
		Age:   54,
	}
	// create a slice of strcture and put the person/s in it
	people := []person{p1, p2, p3}

	fmt.Println(people)

	// pass in the structure to get turned into json.
	// Marshal returns a slice of byte and an error
	bs, err := json.Marshal(people)
	// check the returned error
	if err != nil {
		fmt.Println("error:", err)
	}
	// have to turn the slice of byte into a string for readability
	fmt.Println(string(bs))

}
