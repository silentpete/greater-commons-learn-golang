package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("cemsg: %v", ce.info)
}

func main() {
	ce := customErr{
		info: "this is the custom error info",
	}

	foo(ce)
}

func foo(e error) {
	fmt.Printf("func foo: %v\n", e)
	fmt.Printf("%T\n", e)
	// assertion
	fmt.Println(e.(customErr).info)
}
