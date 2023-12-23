package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Match represents the structure of a football match
type Match struct {
	Competition string `json:"competition"`
	Year        int    `json:"year"`
	Round       string `json:"round"`
	Team1       string `json:"team1"`
	Team2       string `json:"team2"`
	Team1Goals  string `json:"team1goals"`
	Team2Goals  string `json:"team2goals"`
}

// Response represents the structure of the API response
type Response struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Data       []Match `json:"data"`
}

// getNumDraws counts the number of draws in football matches for a given year
func getNumDraws(year int32) int32 {
	apiURL := fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_matches?year=%d", year)

	// Fetch data from the API
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return 0
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return 0
	}

	// Parsing JSON
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return 0
	}

	// Count the number of draws
	numDraws := int32(0)
	for _, match := range response.Data {
		// Assuming a draw is when both teams have the same number of goals
		if match.Team1Goals == match.Team2Goals {
			numDraws++
		}
	}

	return numDraws
}

func main() {
	// Hardcoded API URL for demonstration purposes
	apiURL := "https://jsonmock.hackerrank.com/api/football_matches?year=2011"

	// Fetch data from the API
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Parsing JSON
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Accessing values in the data structure
	fmt.Println("Page:", response.Page)
	fmt.Println("Total Pages:", response.TotalPages)

	// Accessing match data
	for _, match := range response.Data {
		fmt.Printf("\nCompetition: %s\nYear: %d\nRound: %s\n", match.Competition, match.Year, match.Round)
		fmt.Printf("Team 1: %s\nTeam 2: %s\n", match.Team1, match.Team2)
		fmt.Printf("Team 1 Goals: %s\nTeam 2 Goals: %s\n", match.Team1Goals, match.Team2Goals)
	}
}
