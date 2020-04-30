package main

// Prints just the basename of a filepath. Examples below.
// fmt.Println(basename("a/b/c.go")) // "c"
// fmt.Println(basename("c.d.go")) // "c.d"
// fmt.Println(basename("abc")) // "abc"

import (
	"fmt"
	"os"
)

func basename(s string) string {
	// Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
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
