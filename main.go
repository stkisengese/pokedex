package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// clicommand represents a command in the REPL
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// config stores URLs for pagination
type config struct {
	NextURL     string
	PreviousURL string
}

// locationArea represents a location area in the PokeAPI
type locationArea struct {
	Name string `json:"name"`
}

// pokeAPIResponse represents the PokeAPI response
type pokeAPIResponse struct {
	Results  []locationArea `json:"results"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
}

func main() {
	reader := bufio.NewScanner(os.Stdin)

	cfg := &config{}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		if len(words) > 0 {
			commandName := words[0]

			if cmd, found := getCommands()[commandName]; found {
				if err := cmd.callback(cfg); err != nil {
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

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"map": {
			name:        "map",
			description: "Displays 20 location areas in the Pokemon world",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapBack,
		},
	}
}
