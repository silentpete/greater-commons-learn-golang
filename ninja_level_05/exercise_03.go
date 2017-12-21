// Create a new type: vehicle.

// The underlying type is a struct.
// The fields:
// doors
// color
// Create two new types: truck & sedan.

// The underlying type of each of these new types is a struct.
// Embed the “vehicle” type in both truck & sedan.
// Give truck the field “fourWheel” which will be set to bool.
// Give sedan the field “luxury” which will be set to bool. solution
// Using the vehicle, truck, and sedan structs:

// using a composite literal, create a value of type truck and assign values to the fields;
// using a composite literal, create a value of type sedan and assign values to the fields.
// Print out each of these values.

// Print out a single field from each of these values.

package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	fourWheel bool
	vehicle
}
type sedan struct {
	luxury bool
	vehicle
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}
	fmt.Println("s:", s)
	fmt.Println("t:", t)
	fmt.Printf("Type: %T\n", t)
	fmt.Println("Four Wheel:", t.fourWheel)
	fmt.Println("Color:", t.color)
	fmt.Println("Doors:", t.doors)
}
