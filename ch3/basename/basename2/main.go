package main

import (
	"fmt"
	"os"
	"strings"
)

// Prints just the basename of a filepath. Examples below.
// fmt.Println(basename("a/b/c.go")) // "c"
// fmt.Println(basename("c.d.go")) // "c.d"
// fmt.Println(basename("abc")) // "abc"

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	if os.Args[1] == "--" {
		fmt.Println("Basename: " + basename(os.Args[2]))
	} else {
		fmt.Println("Basename: " + basename(os.Args[1]))
	}
}
