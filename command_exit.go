package main

import (
	"fmt"
	"os"

	"github.com/stkisengese/pokedex/internal/models"
)

func commandExit(cfg *models.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
