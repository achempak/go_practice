package main

// Dup2 prints the text of lines that appear more than once
// in the input as well as the files their from.
// It reads from stdin or from a list of named files.

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileDups := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countlines(os.Stdin, "NOFILE", counts, fileDups)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprint(os.Stderr, "dup2: %v\n", err)
				continue // Go to next interation of for loop
			}
			countlines(f, arg, counts, fileDups)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, fileDups[line])
		}
	}
}

func countlines(f *os.File, file_name string, counts map[string]int, fileDups map[string][]string) {
	input := bufio.NewScanner(f)
	// foundDup := false
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 && !foundDup {
			fileDups[input.Text()] = append(fileDups[input.Text()], file_name)
			// foundDup = true
		}
	}
}
