package main

import "fmt"

func commandHelp(commands map[string]cliCommand) func(*config) error {
	return func(*config) error {
		fmt.Println()
		fmt.Println("\nWelcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println()
		for _, cmd := range commands {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}
		fmt.Println()
		return nil
	}
}
