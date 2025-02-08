package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// commandMap handles the 'map' command
func commandMap(cfg *config) error {
	url := cfg.NextURL
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area?limit=20"
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	var response pokeAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}

	if response.Next != nil {
		cfg.NextURL = *response.Next
	} else {
		cfg.NextURL = ""
	}

	if response.Previous != nil {
		cfg.PreviousURL = *response.Previous
	} else {
		cfg.PreviousURL = ""
	}
	return nil
}

// commandMapBack handles the 'mapb' command
func commandMapBack(cfg *config) error {
	if cfg.PreviousURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := http.Get(cfg.PreviousURL)
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	var response pokeAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	if response.Next != nil {
		cfg.NextURL = *response.Next
	} else {
		cfg.NextURL = ""
	}
	if response.Previous != nil {
		cfg.PreviousURL = *response.Previous
	} else {
		cfg.PreviousURL = ""
	}
	return nil
}
