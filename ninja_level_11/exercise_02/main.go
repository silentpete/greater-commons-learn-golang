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

	bs, err := toJSON(p1)
	if err != nil {
		fmt.Println(err)
		// return will exit main... and then the string(bs) will not be executed
		// if using log.Fatalln(err), this would exit, so the return would not be needed.
		return
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	// if there is an error, return the following... with the error
	if err != nil {
		return []byte{}, fmt.Errorf("there was an error: %v", err)
	}
	// if there is NOT an error, return nil
	return bs, nil
}
