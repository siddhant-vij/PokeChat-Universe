package authroutes

import (
	"net/http"
	"net/url"

	"github.com/siddhant-vij/PokeChat-Universe/config"
)

func HandleLogout(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	logoutUrl, err := url.Parse("https://" + cfg.AuthDomain + "/v2/logout")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + r.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := url.Values{}
	params.Add("returnTo", returnTo.String())
	params.Add("client_id", cfg.ClientID)
	logoutUrl.RawQuery = params.Encode()

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		MaxAge:   0,
		Secure:   false,
		HttpOnly: true,
	})

	http.Redirect(w, r, logoutUrl.String(), http.StatusSeeOther)
}
