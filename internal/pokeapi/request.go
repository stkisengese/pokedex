package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/stkisengese/pokedex/internal/models"
)

// handelMapRequest sends an HTTP GET request to the specified URL and processes
// the response. It checks the cache first, and if the data is not found, it makes
// a network request and stores the response in the cache.next and previous URLs if available.
func HandelMapRequest(url string, cfg *models.Config) error {
	if data, ok := cfg.Cache.Get(url); ok {
		fmt.Println("Using cached data")
		return processResponse(data, cfg)
	}

	fmt.Println("cache missed, Making a request...")
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body into a []byte
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	// Add the response to the cache
	cfg.Cache.Add(url, data)

	return processResponse(data, cfg)
}

func processResponse(data []byte, cfg *models.Config) error {
	var response models.PokeAPIResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	// Print the names of the location area
	for _, area := range response.Results {
		fmt.Println(area.Name)
	}

	// Update the configuration with the next and previous URLs
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
