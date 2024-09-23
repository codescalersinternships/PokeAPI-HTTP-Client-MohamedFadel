# PokeAPI CLI

This project implements a command-line interface (CLI) for interacting with the PokeAPI, allowing users to fetch information about Pokemon.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API](#api)
- [Docker Support](#docker-support)
- [Makefile Commands](#makefile-commands)

## Features
- Fetch information about a specific Pokémon by name or ID
- Retrieve a list of Pokémon with pagination support
- Command-line interface for easy interaction
- Logging with Zap logger
- Docker support for easy deployment
- Comprehensive error handling

## Prerequisites
- Go 1.16 or higher
- Docker (optional, for containerization)
- Make (optional, for using Makefile commands)

## Installation
1. Clone the repository:
   ```
   git clone https://github.com/codescalersinternships/PokeAPI-HTTP-Client-MohamedFadel.git
   cd PokeAPI-HTTP-Client-MohamedFadel
   ```
2. Build the application:
   ```
   make build
   ```

## Usage
The PokeAPI CLI supports two main commands:

1. Get information about a specific Pokémon:
   ```
   ./pokeapi-cli get-pokemon --name pikachu
   ```
   or
   ```
   ./pokeapi-cli get-pokemon --name 25
   ```

2. Get a list of Pokémon:
   ```
   ./pokeapi-cli get-pokemons --limit 10 --offset 0
   ```

## API
The CLI interacts with the PokeAPI using the following methods:

### `GetPokemon(idOrName string) (*Pokemon, error)`
Retrieves information about a specific Pokémon by its name or ID.

### `GetPokemons(offset, limit int) (*PokemonList, error)`
Retrieves a list of Pokémon with pagination support.

## Docker Support
The project includes a Dockerfile for containerization. To build and run the Docker image:

```bash
make docker-build
make docker-run
```

## Makefile Commands
The project includes a Makefile with the following commands:

- `make all`: Run all tasks (deps, fmt, lint, test, build)
- `make build`: Build the CLI binary
- `make test`: Run tests for the project
- `make clean`: Clean up build artifacts
- `make run`: Build and run the CLI
- `make deps`: Download dependencies
- `make fmt`: Format the Go code
- `make lint`: Run golangci-lint
- `make docker-build`: Build the Docker image
- `make docker-run`: Run the CLI in a Docker container

To use these commands, run `make <command>` in the project root directory.
