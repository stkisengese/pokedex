package models

import "github.com/stkisengese/pokedex/internal/pokecache"

// clicommand represents a command in the REPL
type CLICommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

// config stores URLs for pagination
type Config struct {
	NextURL     string
	PreviousURL string
	Cache       *pokecache.Cache
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
