package chatroutes

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	stopChannel   = make(chan struct{})
	stopChannelMu sync.Mutex
)

func SseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	clientGone := r.Context().Done()

	loremIpsum := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."
	wordStream := simulateLlmStream(loremIpsum, 50*time.Millisecond)

	for {
		select {
		case <-clientGone:
			fmt.Println("Client disconnected")
			return
		case <-stopChannel:
			fmt.Fprintf(w, "event: Complete\n")
			fmt.Fprintf(w, "data: Stream stopped\n\n")
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			return
		case word, ok := <-wordStream:
			if !ok {
				fmt.Fprintf(w, "event: Complete\n")
				fmt.Fprintf(w, "data: LLM simulation done!\n\n")
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
				return
			}

			fmt.Fprintf(w, "data: %s\n\n", word)

			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		}
	}
}

func StopSseHandler(w http.ResponseWriter, r *http.Request) {
	stopChannelMu.Lock()
	defer stopChannelMu.Unlock()
	close(stopChannel)
	stopChannel = make(chan struct{})
}

func simulateLlmStream(text string, delay time.Duration) <-chan string {
	stream := make(chan string)

	go func() {
		defer close(stream)
		for _, char := range text {
			stream <- string(char)
			time.Sleep(delay)
		}
	}()

	return stream
}
