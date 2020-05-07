package main

import (
	"fmt"
	"os"
)

func joinStrings(strings ...string) string {
	var out string
	for _, s := range strings {
		out += s
	}
	return out
}

func main() {
	first := 1
	if os.Args[1] == "--" {
		first = 2
	}
	fmt.Println(joinStrings(os.Args[first:]...))
}
