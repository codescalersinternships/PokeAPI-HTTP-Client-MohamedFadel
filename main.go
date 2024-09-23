package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/codescalersinternships/PokeAPI-HTTP-Client-MohamedFadel/pokeapi"
	"go.uber.org/zap"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: pokeapi-cli <command> [arguments]")
		fmt.Println("Commands: get-pokemon, get-pokemons")
		os.Exit(1)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	client := pokeapi.NewClient(pokeapi.WithLogger(logger))

	switch os.Args[1] {
	case "get-pokemon":
		getPokemon(client, logger)
	case "get-pokemons":
		getPokemons(client, logger)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func getPokemon(client *pokeapi.Client, logger *zap.Logger) {
	getPokemonCmd := flag.NewFlagSet("get-pokemon", flag.ExitOnError)
	nameOrID := getPokemonCmd.String("name", "", "Name or ID of the Pokemon")

	getPokemonCmd.Parse(os.Args[2:])

	if *nameOrID == "" {
		fmt.Println("Error: --name flag is required")
		getPokemonCmd.PrintDefaults()
		os.Exit(1)
	}

	pokemon, err := client.GetPokemon(*nameOrID)
	if err != nil {
		logger.Error("Failed to get Pokemon", zap.Error(err))
		fmt.Printf("Error: %v\n", err)
		return
	}

	jsonData, err := json.MarshalIndent(pokemon, "", "  ")
	if err != nil {
		logger.Error("Failed to marshal Pokemon data", zap.Error(err))
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(string(jsonData))
}

func getPokemons(client *pokeapi.Client, logger *zap.Logger) {
	getPokemonsCmd := flag.NewFlagSet("get-pokemons", flag.ExitOnError)
	limit := getPokemonsCmd.Int("limit", 20, "Limit the number of Pokemon returned")
	offset := getPokemonsCmd.Int("offset", 0, "Offset for pagination")

	getPokemonsCmd.Parse(os.Args[2:])

	pokemonList, err := client.GetPokemons(*offset, *limit)
	if err != nil {
		logger.Error("Failed to get Pokemon list", zap.Error(err))
		fmt.Printf("Error: %v\n", err)
		return
	}

	jsonData, err := json.MarshalIndent(pokemonList, "", "  ")
	if err != nil {
		logger.Error("Failed to marshal Pokemon list data", zap.Error(err))
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(string(jsonData))
}
