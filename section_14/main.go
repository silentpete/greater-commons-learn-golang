// pointers
// '&var' will give you the pointer address of the variable
// '*var' is a pointer to a var
// '*int' is a Type
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Assign 'a' a value")
	a := 8
	fmt.Println("print pointer 'a'")
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	foo(&a)
	fmt.Println("change value at 'a' pointer location")
	bar(&a)
	fmt.Println("print pointer 'a'")
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

func foo(p *int) {
	fmt.Println(p)
	fmt.Printf("%T\n", p)
}

func bar(p *int) {
	*p = 2
}
