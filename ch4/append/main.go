package main

import "fmt"

// This is how the append function works in Go.
// The "..." makes the function variadic. It accepts any number
// of final arguments.
func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized lienar complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z[:len(x)], x) // a built-in function
	}
	copy(z[len(x):], y) // append elements of y after elements of x
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
	y = appendInt(y, y...) // Append y to itself.
	fmt.Println(y)
}
