// Create a for loop using this syntax
// for condition { }
// Have it print out the years you have been alive.

package main

import (
	"fmt"
	"time"
)

func main() {

	date_born := time.Date(1980, 8, 8, 8, 0, 0, 0, time.UTC)
	// fmt.Println(date_born.Year())

	date_now := time.Now()
	// fmt.Println(date_now.Year())

	for date_born.Before(date_now) {
		fmt.Println(date_now.Year()-date_born.Year(), "years alive")
		break
	}
}

// Wow, talk about making it more complicated than Todd's example, ugh
