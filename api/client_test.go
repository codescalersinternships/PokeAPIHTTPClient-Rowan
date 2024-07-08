package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

func Test_GetPokemonByNameMock(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pokemon := Pokemon{
			Name: "rowan", Height: 170,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&pokemon); err != nil {
			t.Error()
		}
	}))
	defer mockServer.Close()
	c := NewClient()
	ctx := context.Background()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, mockServer.URL, nil)
	if err != nil {
		t.Errorf("error is %v", err)
	}
	request.Header.Add("Accept", "application/json")
	response, err := c.httpClient.Do(request)
	if err != nil {
		t.Errorf("error is %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code returned! %v", response.StatusCode)
	}
	var poke Pokemon
	err = json.NewDecoder(response.Body).Decode(&poke)
	if err != nil {
		t.Errorf("error decoding into json, %v", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, "rowan", poke.Name)
	assert.Equal(t, 170, poke.Height)
}
