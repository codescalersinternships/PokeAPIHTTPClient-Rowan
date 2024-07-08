package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetPokemonByName(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		client := NewClient()
		poke, err := client.GetPokemonByname(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, "pikachu", poke.Name)
	})

}
