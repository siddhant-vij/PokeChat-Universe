package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiUrl = "https://pokeapi.co/api/v2"

var rl = newRateLimiter(50) // 50 requests per second

func do(endpoint string, obj interface{}) error {
	// Block until a token is available from the rate limiter
	rl.wait()

	req, err := http.NewRequest(http.MethodGet, apiUrl+endpoint, nil)
	if err != nil {
		return fmt.Errorf("failed to create client request: %w", err)
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send client request: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(obj)
	if err != nil {
		return fmt.Errorf("failed to decode client response: %w", err)
	}
	return nil
}
