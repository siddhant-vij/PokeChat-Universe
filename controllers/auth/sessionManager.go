package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/siddhant-vij/PokeChat-Universe/config"
)

type UserSession struct {
	SessionId    string
	AccessToken  string
	IpAddress    string
	UserAgent    string
	LastActivity time.Time
	ExpiresAt    time.Time
}

func NewUserSession(cfg *config.AppConfig, accessToken string) (*UserSession, error) {
	sessionId, err := GenerateSessionId()
	if err != nil {
		return nil, fmt.Errorf("failed to create new session id while creating a new user session: %w", err)
	}
	return &UserSession{
		SessionId:    sessionId,
		AccessToken:  accessToken,
		IpAddress:    cfg.IpAddress,
		UserAgent:    cfg.UserAgent,
		LastActivity: time.Now(),
		ExpiresAt:    time.Now().Add(24 * time.Hour),
	}, nil
}

func (s *UserSession) StoreSession(ctx context.Context, cfg *config.AppConfig) error {
	sessionData := map[string]interface{}{
		"access_token":  s.AccessToken,
		"ip_address":    s.IpAddress,
		"user_agent":    s.UserAgent,
		"last_activity": s.LastActivity.Format(time.RFC3339),
		"expires_at":    s.ExpiresAt.Format(time.RFC3339),
	}
	err := cfg.RedisClient.HSet(ctx, "session:"+s.SessionId, sessionData).Err()
	if err != nil {
		return fmt.Errorf("failed to store user session with id: %s in redis. Err: %w", s.SessionId, err)
	}

	cfg.RedisClient.Expire(ctx, "session:"+s.SessionId, time.Until(s.ExpiresAt))
	return nil
}

func GenerateSessionId() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
