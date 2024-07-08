package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetPokemonByName(t *testing.T) {
	t.Run("testing get by name", func(t *testing.T) {
		client := NewClient()
		poke, err := client.GetPokemonByname(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, "pikachu", poke.Name)
	})
	t.Run("testing wrong url", func(t *testing.T) {
		client := NewClient(CustomAPIURL("test"))
		_, err := client.GetPokemonByname(context.Background(), "pikachu")
		assert.Error(t, err)
		assert.Equal(t, "test", client.apiURL)
	})
	t.Run("testing listing all pokemons", func(t *testing.T) {
		client := NewClient()
		poke, err := client.ListAllPokemon(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 100, len(poke.Pokemons))

	})

}
