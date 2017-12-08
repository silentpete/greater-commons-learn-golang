// Create a for loop using this syntax
// for { }
// Have it print out the years you have been alive.

package main

import (
	"fmt"
	"time"
)

func main() {
	date_born := 1980
	date_now := time.Now().Year()
	for {
		if date_born <= date_now {
			fmt.Println(date_born)
			date_born++
		} else {
			break
		}

	}
}
