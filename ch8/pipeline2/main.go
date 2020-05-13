package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	// range syntax here iteratues over naturals until the
	// naturals channel is closed. Same as
	// go func() {
	// 	for {
	// 		x, ok := <-naturals
	// 		if !ok {
	// 			break // channel was closed and drained
	// 		}
	// 		squares <- x * x
	// 		}
	// 		close(squares)
	// 	}()
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
