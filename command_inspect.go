package main

import (
	"errors"
	"fmt"

	"github.com/stkisengese/pokedex/internal/models"
)

// commandInspect handles the 'inspect' command to Display Pokemon details
func commandInspect(cfg *models.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: inspect <pokemon-name>")
	}

	pokemonName := args[0]
	pokemon, exists := cfg.Pokedex[pokemonName]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
