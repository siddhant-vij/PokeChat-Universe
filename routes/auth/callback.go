package authroutes

import (
	"context"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"

	"github.com/siddhant-vij/PokeChat-Universe/cmd/web/templates/pages"
	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/auth"
)

func ServeCallbackPage(w http.ResponseWriter, r *http.Request, authenticator *auth.Authenticator, cfg *config.AppConfig) {
	if r.URL.Query().Get("state") != cfg.SessionState {
		log.Printf("invalid state parameter. Expected: %s, got: %s", cfg.SessionState, r.URL.Query().Get("state"))
		clientErrorPage := pages.ClientErrorPage(cfg.AuthStatus)
		clientErrorPage.Render(r.Context(), w)
		return
	}

	token, err := authenticator.Exchange(
		r.Context(),
		r.URL.Query().Get("code"),
		oauth2.VerifierOption(cfg.PkceCodeVerifier),
	)
	if err != nil {
		log.Printf("failed to convert an authorization code into a token. Err: %v", err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	cfg.AccessTokenIssuedAt = time.Now()

	userDataFromToken, err := authenticator.ExtractUserProfileInfo(cfg, token.AccessToken)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	err = cfg.DBQueries.InsertUser(context.Background(), userDataFromToken)
	if err != nil {
		log.Printf("failed to insert user into DB. Err: %v", err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	cfg.IpAddress = r.RemoteAddr
	cfg.UserAgent = r.UserAgent()
	userSession, err := auth.NewUserSession(cfg, token.AccessToken)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}
	err = userSession.StoreSession(r.Context(), cfg)
	if err != nil {
		log.Println(err)
		serverErrorPage := pages.ServerErrorPage(cfg.AuthStatus)
		serverErrorPage.Render(r.Context(), w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    userSession.SessionId,
		MaxAge:   86400,
		Secure:   false,
		HttpOnly: true,
	})
	cfg.AuthStatus = true
	http.Redirect(w, r, "/pokedex", http.StatusTemporaryRedirect)
}
