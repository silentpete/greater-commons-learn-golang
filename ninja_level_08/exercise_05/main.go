package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byLast []user

func (a byLast) Len() int           { return len(a) }
func (a byLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// your code goes here
	sort.Slice(users, func(i, j int) bool { return users[i].Age < users[j].Age })
	for _, user := range users {
		fmt.Println(user.First, user.Last, "is", user.Age, "years old")
		fmt.Println("and likes to say:")
		sort.Strings(user.Sayings)
		for _, v := range user.Sayings {
			fmt.Println("\t", v)
		}
	}

	sort.Sort(byLast(users))
	for _, user := range users {
		fmt.Println(user.First, user.Last, "is", user.Age, "years old")
		fmt.Println("and likes to say:")
		for _, v := range user.Sayings {
			fmt.Println("\t", v)
		}
	}
}
