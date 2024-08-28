package ui

import (
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/test"
)

func HelloWebHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	postResponse := test.HelloPost(name)
	postResponse.Render(r.Context(), w)
	emptyInput := test.HelloPostOOB()
	emptyInput.Render(r.Context(), w)
}
