package auth

import (
	"context"
	"errors"
	"log"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"

	"github.com/siddhant-vij/PokeChat-Universe/config"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

func NewAuthenticator(cfg *config.AppConfig) *Authenticator {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+cfg.AuthDomain+"/",
	)
	if err != nil {
		log.Fatalf("error creating authenticator. Err: %v", err)
		return nil
	}

	conf := oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.CallbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}
}

func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
