package authroutes

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/auth"
)

func generateSessionId() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func ServeCallbackPage(w http.ResponseWriter, r *http.Request, auth *auth.Authenticator, cfg *config.AppConfig) {
	if r.URL.Query().Get("state") != cfg.SessionState {
		http.Error(w, "Invalid state parameter.", http.StatusBadRequest)
		return
	}

	token, err := auth.Exchange(
		r.Context(),
		r.URL.Query().Get("code"),
		oauth2.VerifierOption(cfg.PkceCodeVerifier),
	)
	if err != nil {
		http.Error(w, "Failed to convert an authorization code into a token.", http.StatusInternalServerError)
		return
	}

	idToken, err := auth.VerifyIDToken(r.Context(), token)
	if err != nil {
		http.Error(w, "Failed to verify ID Token.", http.StatusInternalServerError)
		return
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		http.Error(w, "Failed to parse ID Token claims.", http.StatusInternalServerError)
		return
	}

	sessionId, err := generateSessionId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cfg.SessionTokenMap[sessionId] = token

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		MaxAge:   86400,
		Secure:   false,
		HttpOnly: true,
	})
	http.Redirect(w, r, "/resource", http.StatusTemporaryRedirect)
}
