package api

// Pokemon is our main model, which we parse and manipulate
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

// PokemonArray is an array of Pokemon structs to get parsed later
type PokemonArray struct {
	Pokemons []Pokemon `json:"results"`
}

// NewPokemon is our Pokemon initializer
func NewPokemon() Pokemon {
	return Pokemon{}
}
