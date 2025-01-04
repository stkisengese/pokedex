package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(`Hello, World!`)
}

func cleanInput(text string) []string {
	result := strings.ToLower(text)
	//result = strings.TrimSpace(result)
	return strings.Fields(result)
}
