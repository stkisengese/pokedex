package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		words := cleanInput(input)
		if len(words) > 0 {
			fmt.Printf("Your command was: %s\n", words[0])
		}
	}
}

func cleanInput(text string) []string {
	result := strings.ToLower(text)
	//result = strings.TrimSpace(result)
	return strings.Fields(result)
}
