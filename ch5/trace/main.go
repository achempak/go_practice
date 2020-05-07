// Show use of differ when used to pair "on entry" and "on exit" functions for debugging

package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget extra parentheses
	// ...lots of work...
	time.Sleep(5 * time.Second) // simulate slow operation
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func main() {
	bigSlowOperation()
}
