package models

import "github.com/stkisengese/pokedex/internal/pokecache"

// clicommand represents a command in the REPL
type CLICommand struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}

// config stores URLs for pagination
type Config struct {
	NextURL     string
	PreviousURL string
	Cache       *pokecache.Cache
	Pokedex     map[string]Pokemon
}

// locationArea represents a location area in the PokeAPI
type locationArea struct {
	Name string `json:"name"`
}

// pokeAPIResponse represents the PokeAPI response
type PokeAPIResponse struct {
	Results  []locationArea `json:"results"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
}

type LocationAreaDetail struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
