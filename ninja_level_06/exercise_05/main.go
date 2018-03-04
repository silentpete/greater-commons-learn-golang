// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area = π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

package main

import "fmt"

func main() {
	c1 := circle{
		diameter: 24.69,
	}
	s1 := square{
		sideLength: 15,
	}
	// c1.area()
	info(c1)
	// s1.area()
	info(s1)
}

type square struct {
	sideLength float32
}
type circle struct {
	diameter float32
}

func (s square) area() {
	sq := s.sideLength * s.sideLength
	fmt.Println("area of square is:", sq, "feet^2")
}
func (c circle) area() {
	r := c.diameter / 2
	r2 := r * r
	pieRSq := r2 * 3.14
	fmt.Println("area of circle is:", pieRSq, "feet^2")
	// fmt.Println(pieRSq)
}

type shape interface {
	area()
}

func info(s shape) {
	s.area()
}
