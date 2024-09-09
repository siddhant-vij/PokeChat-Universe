package routes

import (
	"fmt"
	"net/http"
)

func AvailableRedirectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Redirect", "/pokedex")
}

func CollectedRedirectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Redirect", "/collectedPokedex")
}

func ChatRedirectHandler(w http.ResponseWriter, r *http.Request) {
	pokemonName := r.FormValue("pokemonName")
	w.Header().Set("HX-Redirect", fmt.Sprintf("/chat/%s", pokemonName))
}
