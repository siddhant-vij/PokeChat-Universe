package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/controllers/auth"
)

func IsAuthenticated(next http.Handler, cfg *config.AppConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionIdCookie, err := r.Cookie("session_id")
		if err != nil {
			log.Printf("error getting the session_id cookie. Err: %v", err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		err = VerifySessionBinding(r.Context(), sessionIdCookie.Value, cfg)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		sessionId, err := RegenerateUserSessionIfNeeded(r.Context(), sessionIdCookie.Value, cfg)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    sessionId,
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
			MaxAge:   int(time.Until(cfg.AccessTokenIssuedAt.Add(24 * time.Hour)).Seconds()),
		})

		next.ServeHTTP(w, r)
	})
}

func VerifySessionBinding(ctx context.Context, sessionId string, cfg *config.AppConfig) error {
	storedIpAddress, err := cfg.RedisClient.HGet(ctx, "session:"+sessionId, "ip_address").Result()
	if err != nil {
		return fmt.Errorf("failed to get stored ip_address for session %s from redis: %w", sessionId, err)
	}

	storedUserAgent, err := cfg.RedisClient.HGet(ctx, "session:"+sessionId, "user_agent").Result()
	if err != nil {
		return fmt.Errorf("failed to get stored user_agent for session %s from redis: %w", sessionId, err)
	}

	if storedIpAddress != cfg.IpAddress || storedUserAgent != cfg.UserAgent {
		return fmt.Errorf("suspicious session activity. session %s is no longer valid", sessionId)
	}
	return nil
}

func RegenerateUserSessionIfNeeded(ctx context.Context, sessionId string, cfg *config.AppConfig) (string, error) {
	if time.Since(cfg.AccessTokenIssuedAt) >= 24*time.Hour {
		return sessionId, fmt.Errorf("access token for session %s has expired", sessionId)
	}

	sessionData, err := cfg.RedisClient.HGetAll(ctx, "session:"+sessionId).Result()
	if err != nil {
		return sessionId, fmt.Errorf("failed to get session data for session %s from redis: %w", sessionId, err)
	}

	lastActivity, err := time.Parse(time.RFC3339, sessionData["last_activity"])
	if err != nil {
		return sessionId, fmt.Errorf("failed to parse last_activity for session %s: %w", sessionId, err)
	}

	if time.Since(lastActivity) > 30*time.Minute {
		newSessionId, err := auth.GenerateSessionId()
		if err != nil {
			return sessionId, fmt.Errorf("failed to create new session id while regenerating user session: %w", err)
		}

		pipe := cfg.RedisClient.TxPipeline()

		pipe.HSet(ctx, "session:"+newSessionId, sessionData)

		pipe.HSet(ctx, "session:"+newSessionId, "last_activity", time.Now().Format(time.RFC3339))

		pipe.Expire(ctx, "session:"+newSessionId, time.Until(cfg.AccessTokenIssuedAt.Add(24*time.Hour)))

		pipe.Del(ctx, "session:"+sessionId)

		if _, err := pipe.Exec(ctx); err != nil {
			return sessionId, fmt.Errorf("failed to execute redis transaction pipeline: %w", err)
		}
		return newSessionId, nil
	}
	return sessionId, nil
}
