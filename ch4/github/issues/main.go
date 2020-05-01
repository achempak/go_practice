package main

import (
	"fmt"
	"log"
	"os"

	"goPractice/ch4/github/github"
)

func main() {
	switch requestType := os.Args[1]; requestType {
	case "list":
		result, err := github.SearchIssues(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d issues:\n", result.TotalCount)
		for _, item := range result.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	case "get":
		result, err := github.SearchIssues(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf
	}
}
