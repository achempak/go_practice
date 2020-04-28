package main

// Reverse a slice of ints in place

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Shift left or right by n spots
func shift(s []int, n int, left bool) {
	if left {
		reverse(s[:n])
		reverse(s[n:])
		reverse(s)
	} else {
		reverse(s)
		reverse(s[:n])
		reverse(s[n:])
	}
}

// Auxiliary function to compute greatest common divisor.
// Uses the Euclidean algorithm. Ideally larger number is x and
// smaller number is y.
func gcd(x int, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

// "Leapfrog method." We need the GCD of the length of the array and the
// step size in order to detect if we've reached the point we started at,
// so that we can break from the loop and move to the next set of elements
// to leap over.
func shift1Pass(s []int, n int, left bool) {
	d := gcd(len(s), n)
	if left {
		for i := 0; i < d; i++ {
			temp := s[i] // Store first value since it'll be overwritten
			j := i
			for {
				k := j + n
				if k >= len(s) {
					k = k - len(s)
				}
				if k == i {
					break
				}
				s[j] = s[k]
				j = k
			}
			s[j] = temp
		}
	} else {
		for i := 0; i < d; i++ {
			temp := s[i]
			j := i
			for {
				k := j - n
				if k < 0 {
					k = k + len(s)
				}
				if k == i {
					break
				}
				s[j] = s[k]
				j = k
			}
			s[j] = temp
		}
	}
}

func main() {
	shiftVal := 4
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Printf("Original array: %v\n", s)
	shift(s, shiftVal, true)
	fmt.Printf("Shifted left %d: %v\n", shiftVal, s)
	shift(s, shiftVal, false)
	fmt.Printf("Shift right %d to get back to original: %v\n\n", shiftVal, s)

	s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Printf("Original array: %v\n", s)
	shift1Pass(s, shiftVal, true)
	fmt.Printf("Shifted left %d: %v\n", shiftVal, s)
	shift1Pass(s, shiftVal, false)
	fmt.Printf("Shift right %d to get back to original: %v\n", shiftVal, s)
}
