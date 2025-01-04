package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// commandMap handles the 'map' command
func commandMap(cfg *config) error {
	url := cfg.NextURL
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area?limit=20"
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	var response pokeAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}

	if response.Next != nil {
		cfg.NextURL = *response.Next
	} else {
		cfg.NextURL = ""
	}

	if response.Previous != nil {
		cfg.PreviousURL = *response.Previous
	} else {
		cfg.PreviousURL = ""
	}
	return nil
}

// commandMapBack handles the 'mapb' command
func commandMapBack(cfg *config) error {
	if cfg.PreviousURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := http.Get(cfg.PreviousURL)
	if err != nil {
		return fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	var response pokeAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	for _, area := range response.Results {
		fmt.Println(area.Name)
	}
	if response.Next != nil {
		cfg.NextURL = *response.Next
	} else {
		cfg.NextURL = ""
	}
	if response.Previous != nil {
		cfg.PreviousURL = *response.Previous
	} else {
		cfg.PreviousURL = ""
	}
	return nil
}

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

func main() {
	reader := bufio.NewScanner(os.Stdin)

	cfg := &config{}

	// Create a map of supported commands
	commands := map[string]cliCommand{}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp(commands),
	}

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays 20 location areas in the Pokemon world",
		callback:    commandMap,
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous 20 location areas",
		callback:    commandMapBack,
	}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())

		if len(words) > 0 {
			commandName := words[0]
			if cmd, found := commands[commandName]; found {
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
