package middlewares

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/auth"
)

func IsAuthenticated(next http.Handler, cfg *config.AppConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionIdCookie, err := r.Cookie("session_id")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		err = VerifySessionBinding(r.Context(), sessionIdCookie.Value, cfg)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		sessionId, err := RegenerateUserSessionIfNeeded(r.Context(), sessionIdCookie.Value, cfg)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionId,
			Secure:   false,
			HttpOnly: true,
			MaxAge:   int(time.Until(cfg.AccessTokenIssuedAt.Add(24 * time.Hour)).Seconds()),
			SameSite: http.SameSiteStrictMode,
		})

		next.ServeHTTP(w, r)
	})
}

func VerifySessionBinding(ctx context.Context, sessionId string, cfg *config.AppConfig) error {
	storedIpAddress, err := cfg.RedisClient.HGet(ctx, "session:"+sessionId, "ip_address").Result()
	if err != nil {
		return err
	}

	storedUserAgent, err := cfg.RedisClient.HGet(ctx, "session:"+sessionId, "user_agent").Result()
	if err != nil {
		return err
	}

	if storedIpAddress != cfg.IpAddress || storedUserAgent != cfg.UserAgent {
		return errors.New("suspicious activity detected")
	}
	return nil
}

func RegenerateUserSessionIfNeeded(ctx context.Context, sessionId string, cfg *config.AppConfig) (string, error) {
	if time.Since(cfg.AccessTokenIssuedAt) >= 24*time.Hour {
		return sessionId, errors.New("access token expired")
	}

	sessionData, err := cfg.RedisClient.HGetAll(ctx, "session:"+sessionId).Result()
	if err != nil {
		return sessionId, err
	}

	lastActivity, _ := time.Parse(time.RFC3339, sessionData["last_activity"])
	if time.Since(lastActivity) > 30*time.Minute {
		newSessionId, _ := auth.GenerateSessionId()

		pipe := cfg.RedisClient.TxPipeline()

		pipe.HSet(ctx, "session:"+newSessionId, sessionData)

		pipe.HSet(ctx, "session:"+newSessionId, "last_activity", time.Now().Format(time.RFC3339))

		pipe.Expire(ctx, "session:"+newSessionId, time.Until(cfg.AccessTokenIssuedAt.Add(24*time.Hour)))

		pipe.Del(ctx, "session:"+sessionId)

		if _, err := pipe.Exec(ctx); err != nil {
			return sessionId, err
		}
		return newSessionId, nil
	}
	return sessionId, nil
}
