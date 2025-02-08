package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/stkisengese/pokedex/internal/models"
	"github.com/stkisengese/pokedex/internal/pokecache"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	// Initialize the cache with a 5-minute expiration interval
	cache := pokecache.NewCache(5 * time.Minute)

	cfg := &models.Config{
		Cache: cache,
	}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		if len(words) > 0 {
			commandName := words[0]
			args := words[1:]

			if cmd, found := getCommands()[commandName]; found {
				if err := cmd.Callback(cfg, args...); err != nil {
					fmt.Printf("Error executing command '%s': %s\n", commandName, err)
				}
				continue
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}

	}
}

func cleanInput(text string) []string {
	result := strings.ToLower(text)
	return strings.Fields(result)
}

func getCommands() map[string]models.CLICommand {
	return map[string]models.CLICommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},

		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},

		"map": {
			Name:        "map",
			Description: "Displays 20 location areas in the Pokemon world",
			Callback:    commandMap,
		},

		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 location areas",
			Callback:    commandMapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore list of all the Pokemon in a location area",
			Callback:    commandExplore,
		},
	}
}
