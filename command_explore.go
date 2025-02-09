package main

import (
	"encoding/json"
	"fmt"

	"github.com/stkisengese/pokedex/internal/models"
	"github.com/stkisengese/pokedex/internal/pokeapi"
)

func commandExplore(cfg *models.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <location-area-name>")
	}

	locationAreaName := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationAreaName)

	// Fetch data from the cache or make a new request
	data, err := pokeapi.FetchData(url, cfg.Cache)
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}

	// Parse the response
	var locationArea models.LocationAreaDetail
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
