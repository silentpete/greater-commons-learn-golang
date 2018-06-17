package main

import (
	"fmt"

	"./dog"
)

func main() {
	dogAge := dog.Years(3)
	fmt.Println(dogAge)
}
