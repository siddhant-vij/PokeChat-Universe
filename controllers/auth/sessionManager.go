package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
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

func NewUserSession(cfg *config.AppConfig, accessToken string) *UserSession {
	sessionId, _ := GenerateSessionId()
	return &UserSession{
		SessionId:    sessionId,
		AccessToken:  accessToken,
		IpAddress:    cfg.IpAddress,
		UserAgent:    cfg.UserAgent,
		LastActivity: time.Now(),
		ExpiresAt:    time.Now().Add(24 * time.Hour),
	}
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
		return err
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
