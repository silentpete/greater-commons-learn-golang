// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

package main

import (
	"fmt"
)

func main() {
	favSport := "climbing"
	switch favSport {
	case "hiking":
		fmt.Println("hiking")
	case "biking":
		fmt.Println("biking")
	case "climbing":
		fmt.Println("climbing")
	default:
		fmt.Println("nothing")
	}

}
