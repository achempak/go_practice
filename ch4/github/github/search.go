package github

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

//OAuth token
func token(file string) string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("%v\n", err)
		return "NO_TOKEN"
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return scanner.Text()
}

// SearchIssues queries GitHub issue tracker
// example. repo:golang/go is:open json decoder
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Println("URL: " + IssuesURL + q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// GetIssue gets a specific GitHub issue.
// GET /repos/:owner/:repo/issues/:issue_number
func GetIssue(terms []string) (*Issue, error) {
	q := strings.Join(terms, "/")
	fmt.Println("URL: " + IssueURL + q)
	resp, err := http.Get(IssueURL + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Search query failed: %s", resp.Status)
	}
	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// CreateIssue posts a GitHub issue.
// POST /repos/:owner/:repo/issues
func CreateIssue(terms []string, title string) (*http.Response, error) {
	q := strings.Join(terms, "/")
	fmt.Println("URL: " + IssueURL + q)

	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return nil, fmt.Errorf("Couldn't create temporary file to write issue body")
	}
	filename := file.Name()
	// Defere removal of temporary file in case any fo the next steps fails
	defer os.Remove(filename)
	if err := file.Close(); err != nil {
		return nil, err
	}
	cmd := exec.Command("vim", filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	body := string(data)

	post, err := json.Marshal(Issue{Title: title, Body: body})
	if err != nil {
		return nil, err
	}
	fmt.Println(string(post))

	req, err := http.NewRequest("POST", IssueURL+q, bytes.NewBuffer(post))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "token "+token("./token.txt"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("Post failed: %s", resp.Status)
	}
	resp.Body.Close()
	return resp, nil
}
