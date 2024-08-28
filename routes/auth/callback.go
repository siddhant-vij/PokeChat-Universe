package authroutes

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/oauth2"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/auth"
)

func ServeCallbackPage(w http.ResponseWriter, r *http.Request, authenticator *auth.Authenticator, cfg *config.AppConfig) {
	if r.URL.Query().Get("state") != cfg.SessionState {
		http.Error(w, "Invalid state parameter.", http.StatusBadRequest)
		return
	}

	token, err := authenticator.Exchange(
		r.Context(),
		r.URL.Query().Get("code"),
		oauth2.VerifierOption(cfg.PkceCodeVerifier),
	)
	if err != nil {
		http.Error(w, "Failed to convert an authorization code into a token.", http.StatusInternalServerError)
		return
	}

	cfg.AccessTokenIssuedAt = time.Now()

	userDataFromToken, err := authenticator.ExtractUserProfileInfo(cfg, token.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = cfg.DBQueries.InsertUser(context.Background(), userDataFromToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cfg.IpAddress = r.RemoteAddr
	cfg.UserAgent = r.UserAgent()
	userSession := auth.NewUserSession(cfg, token.AccessToken)
	err = userSession.StoreSession(r.Context(), cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    userSession.SessionId,
		MaxAge:   86400,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	http.Redirect(w, r, "/resource", http.StatusTemporaryRedirect)
}
