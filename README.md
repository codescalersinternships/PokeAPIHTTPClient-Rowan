# PokeAPIHTTPClient-Rowan
Creating an HTTP client in Go that consumes the  Pokémon APIs.

## 1. Client Functionality:

- `Get Pokemon`
- `List All Pokemon`

## 2. Implementations:

The client is implemented using:
1. Standard library `(net/http)`

## 3. Directories Walkthrough
        root  
        |____cli
        |       |___main.go
        |                      
        |
        |____api
                |___client.go (pkg code)
                |___client_test.go (test file)

## 5. Package Testing
### How Testing Is Done (Integration/Mock):
- Using `httptest` package, we can mock a server, so that our client sends actual requests but to a mocked server.
- By accessing mockServer.URL --> we can send its path to our client smoothly.
- Tests include:
  - Unit tests (calling api directly)
  - Integration mock tests (mocking the pokeapi)

### Coverage
- Test coverage: 85.1% of statements

## 6. Package Importing
```
import
"github.com/codescalersinternships/PokeAPIHTTPClient-Rowan/api"
```
If not available
```
go get github.com/codescalersinternships/PokeAPIHTTPClient-Rowan/api
```
## 7. UseCases

<!-- ### Commands:  -->
 To use package, user must intialize it using a server URL
```
client := api.NewClient()
```
 User can get a pokemon by specifying its name
```
poke, err := client.GetPokemonByname(context.Background(), "pikachu")
```
User can get a list of pokemons by invoking List All function
```
pokemons, err := client.ListAllPokemon(context.Background())
```
