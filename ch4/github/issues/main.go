package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

		// Sort results: less than 182 days, between 182-365 days, greater than 365 days.
		// Use Dutch flag sorting algorithm.
		low, mid, high := 0, 0, len(result.Items)-1
		for i := 0; mid <= high; i++ {
			duration := time.Since(result.Items[i].CreatedAt).Hours() / 24
			if duration < 182 {
				result.Items[low], result.Items[i] = result.Items[i], result.Items[low]
				low, mid = low+1, mid+1
			} else if duration >= 182 && duration < 365 {
				mid++
			} else {
				result.Items[high], result.Items[i] = result.Items[i], result.Items[high]
				i, high = i-1, high-1
			}
		}

		// %-5d is left justified with 5 digits of precision
		// %-9.9s is left justified string, min and max number of chars is 9
		// %55.55s is left justified string, min and max number of chars is 55
		fmt.Println("Issues less than 6 months old:")
		for i := 0; i < low; i++ {
			item := result.Items[i]
			fmt.Printf("#%-5d %-9.9s %-55.55s %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
		fmt.Println("\nIssues between 6 months and one year old:")
		for i := low; i < mid; i++ {
			item := result.Items[i]
			fmt.Printf("#%-5d %-9.9s %-55.55s %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
		}
		fmt.Println("\nIssues older than one year:")
		for i := mid; i < len(result.Items); i++ {
			item := result.Items[i]
			fmt.Printf("#%-5d %-9.9s %-55.55s %v\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt)
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
