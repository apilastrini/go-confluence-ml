package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Page represents a Confluence page
type Page struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"body"`
}

func main() {
	// Replace 'ATLASSIAN_TOKEN' with your actual Atlassian API token
	token := ""

	// Construct HTTP client with token in header
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://vxfiber.atlassian.net/wiki/rest/api/content", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	// Make HTTP request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Print JSON response...just for debug!
	fmt.Println("JSON Response:", string(body))

	// Unmarshal JSON response
	var page Page
	err = json.Unmarshal(body, &page)
	if err != nil {
		panic(err)
	}

	// Print page title
	fmt.Println("Page Title:", page.Title)
}
