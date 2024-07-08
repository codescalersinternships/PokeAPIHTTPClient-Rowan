package main

import (
	"context"
	"fmt"
	"github.com/codescalersinternships/PokeAPIHTTPClient-Rowan/api"
)

func main() {
	client := api.NewClient()
	var poke api.Pokemon
	poke, _ = client.GetPokemonByname(context.Background(), "pikachu")
	fmt.Println(poke.BaseExperience)
	fmt.Println(poke)
	// var pokemons api.PokemonArray
	// pokemons, _ = client.ListAllPokemon(context.Background())
	// fmt.Println(err)
	// for _, poke := range pokemons.Pokemons {
	// 	fmt.Println(poke.Name)
	// }
	// fmt.Println(pokemons)

}
