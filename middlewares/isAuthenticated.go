package middlewares

import (
	"net/http"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/auth"
)

func IsAuthenticated(next http.Handler, auth *auth.Authenticator, cfg *config.AppConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionId, err := r.Cookie("session_id")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		_, ok := cfg.SessionTokenMap[sessionId.Value]
		if !ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
