package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/oauth2"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/database"
)

type Authenticator struct {
	oauth2.Config
}

type UserInfoData struct {
	AuthId     string `json:"sub"`
	Username   string `json:"name"`
	Email      string `json:"email"`
	PictureUrl string `json:"picture"`
}

func NewAuthenticator(cfg *config.AppConfig) *Authenticator {
	endpoint := oauth2.Endpoint{
		AuthURL:       fmt.Sprintf("https://%s/authorize", cfg.AuthDomain),
		DeviceAuthURL: fmt.Sprintf("https://%s/oauth/device/code", cfg.AuthDomain),
		TokenURL:      fmt.Sprintf("https://%s/oauth/token", cfg.AuthDomain),
	}

	conf := oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.CallbackURL,
		Endpoint:     endpoint,
		Scopes:       []string{"openid", "profile", "email"},
	}

	return &Authenticator{
		Config: conf,
	}
}

func (a *Authenticator) ExtractUserProfileInfo(cfg *config.AppConfig, accessToken string) (database.InsertUserParams, error) {
	data := &UserInfoData{}
	err := do(cfg.AuthDomain, "userinfo", accessToken, data)
	if err != nil {
		return database.InsertUserParams{}, err
	}

	var insertUserParams = database.InsertUserParams{
		ID:         uuid.New(),
		AuthID:     data.AuthId,
		Username:   extractName(data),
		Email:      data.Email,
		PictureUrl: data.PictureUrl,
	}
	return insertUserParams, nil
}

func do(baseurl, endpoint, accessToken string, data *UserInfoData) error {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://%s/%s", baseurl, endpoint),
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}

func extractName(userInfo *UserInfoData) string {
	if strings.Contains(userInfo.AuthId, "auth0") {
		// For username & password login
		nickname := strings.Split(userInfo.Username, "@")[0]
		return fmt.Sprintf("%s%s", strings.ToTitle(nickname[0:1]), nickname[1:])
	} else {
		// Valid for google login
		return userInfo.Username
	}
}
