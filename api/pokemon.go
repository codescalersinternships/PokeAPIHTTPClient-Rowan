package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetPokemonByname is a method of Client struct, used to fetch pokemon by its name field
func (c *Client) GetPokemonByname(ctx context.Context, pokemonName string) (Pokemon, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.apiURL+"/api/v2/pokemon/"+pokemonName, nil)
	if err != nil {
		return NewPokemon(), err
	}
	request.Header.Add("Accept", "application/json")
	response, err := c.httpClient.Do(request)
	if err != nil {
		return NewPokemon(), err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return NewPokemon(), fmt.Errorf("unexpected status code returned! %v", response.StatusCode)
	}
	var poke Pokemon
	err = json.NewDecoder(response.Body).Decode(&poke)
	if err != nil {
		return NewPokemon(), fmt.Errorf("error decoding into json, %v", err)
	}
	return poke, nil
}

// ListAllPokemon is a method of Client struct to list all pokemons
func (c *Client) ListAllPokemon(ctx context.Context) (PokemonArray, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.apiURL+"/api/v2/pokemon?limit=100", nil)
	if err != nil {
		return PokemonArray{}, err
	}
	request.Header.Add("Accept", "application/json")
	response, err := c.httpClient.Do(request)
	fmt.Println(response.Body)
	if err != nil {
		return PokemonArray{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return PokemonArray{}, fmt.Errorf("unexpected status code returned! %v", response.StatusCode)
	}
	var poke PokemonArray

	err = json.NewDecoder(response.Body).Decode(&poke)
	if err != nil {
		return PokemonArray{}, fmt.Errorf("error decoding into json, %v", err)
	}
	return poke, nil
}