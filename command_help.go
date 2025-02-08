package main

import (
	"fmt"

	"github.com/stkisengese/pokedex/internal/models"
)

func commandHelp(cfg *models.Config, args ...string) error {
	fmt.Println()
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}
