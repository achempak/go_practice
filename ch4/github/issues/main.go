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
			// %-5d is left justified with 5 digits of precision
			// %9.9s is right justified string, min and max number of chars is 9
			// %.55s is right justified string, max number of chars is 55
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	case "get":
		result, err := github.GetIssue(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("#%-5d, %9.9s %.55s\n%s",
			result.Number, result.User.Login, result.Title, result.Body)
	case "create":
		// First command argument is "create", second is title of issue
		resp, err := github.CreateIssue(os.Args[3:], os.Args[2])
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Success! Response:\n%+v", *resp)
		}
	}
}
