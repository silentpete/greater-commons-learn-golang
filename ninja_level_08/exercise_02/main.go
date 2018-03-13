package main

import (
	"encoding/json"
	"fmt"
)

type jsonMap struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	// the json data
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	// turn the json data into a slice of byte
	bs := []byte(s)
	// create a new data structure empty object
	var stuff []jsonMap
	// Unmarshal the json data into the json structure at the memory location
	err := json.Unmarshal(bs, &stuff)
	// check the return, which unmarshal only returns an error
	if err != nil {
		fmt.Println(err)
	}
	// can range over 'stuff' to print out contents
	for i, v := range stuff {
		fmt.Println("index =", i)
		fmt.Println("\t", v.First, v.Last, "says:")
		for _, n := range v.Sayings {
			fmt.Println("\t -", n)
		}
	}

}
