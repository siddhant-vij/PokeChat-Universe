package client

import (
	"time"
)

type rateLimiter struct {
	tokens chan struct{}
	rate   time.Duration
}

func newRateLimiter(requestsPerSecond int) *rateLimiter {
	rl := &rateLimiter{
		tokens: make(chan struct{}, requestsPerSecond),
		rate:   time.Second / time.Duration(requestsPerSecond),
	}

	go rl.fillTokens()

	return rl
}

func (rl *rateLimiter) fillTokens() {
	ticker := time.NewTicker(rl.rate)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case rl.tokens <- struct{}{}:
		default:
			// If the buffer is full, don't block
		}
	}
}

func (rl *rateLimiter) wait() {
	<-rl.tokens
}
