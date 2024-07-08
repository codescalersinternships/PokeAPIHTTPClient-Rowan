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

// NewPokemon is our Pokemon initializer
func NewPokemon() Pokemon {
	return Pokemon{}
}
