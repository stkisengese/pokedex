package main

import (
	"fmt"

	"github.com/stkisengese/pokedex/internal/models"
)

func commandPokedex(cfg *models.Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for key := range cfg.Pokedex {
		fmt.Printf("  - %s\n", key)
	}
	return nil
}
