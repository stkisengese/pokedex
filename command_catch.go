package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/stkisengese/pokedex/internal/models"
	"github.com/stkisengese/pokedex/internal/pokeapi"
)

// commandCatch handles the 'catch' command
func commandCatch(cfg *models.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: catch <pokemon-name>")
	}

	pokemonName := args[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)

	// Fetch data from the cache or make a new request
	data, err := pokeapi.FetchData(url, cfg.Cache)
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}

	// Parse the response
	var pokemon models.Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	// Simulate catching the Pokémon
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	catchProbability := calculateCatchProbability(pokemon.BaseExperience)
	if rand.Float64() < catchProbability {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.Pokedex[pokemonName] = pokemon // Add to Pokedex
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

// calculateCatchProbability calculates the catch probability based on base experience
func calculateCatchProbability(baseExperience int) float64 {
	// Higher base experience means lower catch probability
	return 1.0 / (1.0 + float64(baseExperience)/100.0)
}
