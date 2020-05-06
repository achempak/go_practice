package main

import (
	"fmt"
)

// prereqs maps computer science courses to the prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math", "operating systems"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	// onStack keeps track of items currently on the stack. Used to check for cycles.
	onStack := make(map[string]bool)
	var foundCycle bool
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				onStack[item] = true
				visitAll(m[item])
				if foundCycle {
					order = append(order, item+" (cycle detected)")
					foundCycle = false
				} else {
					order = append(order, item)
				}
				onStack[item] = false
			} else if onStack[item] {
				foundCycle = true
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	//sort.Strings(keys)
	visitAll(keys)
	return order
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
