package pokeapi

import (
	"fmt"
	"io"
	"net/http"

	"github.com/stkisengese/pokedex/internal/pokecache"
)

func FetchData(url string, cache *pokecache.Cache) ([]byte, error) {
	// Check the cache first
	if data, ok := cache.Get(url); ok {
		fmt.Println("Cache hit! Using cached data.")
		return data, nil
	}

	// If not in the cache, make a network request
	fmt.Println("Cache miss! Making a request...")
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Add the response to the cache
	cache.Add(url, data)
	return data, nil
}
