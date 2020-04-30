package main

import "fmt"

// Example of anonymous fields and embedded structs

type Point struct {
	X, Y int
}

type Circle struct {
	Point  // Equivalent to Center Point, but now anonymous for easier field access
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	// w = Wheel{Circle{Point{8,8}, 5}, 20} this is equivalent to above

	fmt.Printf("%#v\n", w)
	w.X = 42
	fmt.Printf("%#v\n", w)
}
