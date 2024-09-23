package pokeapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestGetPokemon(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v2/pokemon/pikachu" {
			t.Errorf("Expected to request '/api/v2/pokemon/pikachu', got: %s", r.URL.Path)
		}
		if r.Header.Get("User-Agent") != "Test-Agent" {
			t.Errorf("Expected User-Agent header to be 'Test-Agent', got: %s", r.Header.Get("User-Agent"))
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Pokemon{
			ID:             25,
			Name:           "pikachu",
			BaseExperience: 112,
			Height:         4,
			Weight:         60,
		})
	}))
	defer ts.Close()

	logger, _ := zap.NewDevelopment()
	client := NewClient(
		WithBaseURL(ts.URL+"/api/v2"),
		WithTimeout(5*time.Second),
		WithUserAgent("Test-Agent"),
		WithLogger(logger),
	)

	pokemon, err := client.GetPokemon("pikachu")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if pokemon == nil {
		t.Fatal("Expected pokemon to not be nil")
	}
	if pokemon.ID != 25 {
		t.Errorf("Expected Pokemon ID to be 25, got %d", pokemon.ID)
	}
	if pokemon.Name != "pikachu" {
		t.Errorf("Expected Pokemon name to be 'pikachu', got %s", pokemon.Name)
	}
	if pokemon.BaseExperience != 112 {
		t.Errorf("Expected Pokemon base experience to be 112, got %d", pokemon.BaseExperience)
	}
	if pokemon.Height != 4 {
		t.Errorf("Expected Pokemon height to be 4, got %d", pokemon.Height)
	}
	if pokemon.Weight != 60 {
		t.Errorf("Expected Pokemon weight to be 60, got %d", pokemon.Weight)
	}
}

func TestGetPokemonError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	logger, _ := zap.NewDevelopment()
	client := NewClient(
		WithBaseURL(ts.URL+"/api/v2"),
		WithTimeout(5*time.Second),
		WithLogger(logger),
	)

	pokemon, err := client.GetPokemon("nonexistent")
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	if pokemon != nil {
		t.Fatal("Expected pokemon to be nil")
	}
	if err.Error() != "unexpected status code: 404" {
		t.Errorf("Expected error message 'unexpected status code: 404', got '%s'", err.Error())
	}
}
