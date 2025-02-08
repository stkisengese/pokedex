package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/stkisengese/pokedex/internal/models"
	"github.com/stkisengese/pokedex/internal/pokecache"
)

type LocationAreaDetail struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *models.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <location-area-name>")
	}

	locationAreaName := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationAreaName)

	// Fetch data from the cache or make a new request
	data, err := FetchData(url, cfg.Cache)
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}

	// Parse the response
	var locationArea LocationAreaDetail
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	// Display the Pokémon names
	fmt.Printf("Exploring %s...\n", locationAreaName)
	fmt.Println("Found Pokémon:")
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func FetchData(url string, cache *pokecache.Cache) ([]byte, error) {
	// Check the cache first
	if data, ok := cache.Get(url); ok {
		fmt.Println("Cache hit! Using cached data.")
		return data, nil
	}

	// If not in the cache, make a network request
	fmt.Println("Cache miss! Making a request...")
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Add the response to the cache
	cache.Add(url, data)
	return data, nil
}
