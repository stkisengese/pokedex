package main

import (
	"fmt"

	"github.com/stkisengese/pokedex/internal/models"
	"github.com/stkisengese/pokedex/internal/pokeapi"
)

// commandMap handles the 'map' command
func commandMap(cfg *models.Config, args ...string) error {
	url := cfg.NextURL
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area?limit=20"
	}
	return pokeapi.HandelMapRequest(url, cfg)
}

// commandMapBack handles the 'mapb' command
func commandMapBack(cfg *models.Config, args ...string) error {
	if cfg.PreviousURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	return pokeapi.HandelMapRequest(cfg.PreviousURL, cfg)
}
