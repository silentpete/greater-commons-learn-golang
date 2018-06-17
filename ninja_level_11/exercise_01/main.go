package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("on my", err)
	}
	fmt.Println(string(bs))

	// just playing
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	for _, v := range p1.Sayings {
		fmt.Println(v)
	}

}
