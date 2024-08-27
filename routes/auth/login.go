package authroutes

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/auth"
)

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func ServeLoginPage(w http.ResponseWriter, r *http.Request, auth *auth.Authenticator, cfg *config.AppConfig) {
	state, err := generateRandomState()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cfg.SessionState = state
	http.Redirect(w, r,
		auth.AuthCodeURL(
			state,
			oauth2.S256ChallengeOption(cfg.PkceCodeVerifier),
		),
		http.StatusTemporaryRedirect)
}
