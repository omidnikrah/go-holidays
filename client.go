package holidays

import (
	"net/http"
	"sync"
	"time"
)

type client struct {
	httpClient *http.Client
	cache      map[string][]Holiday
	cacheTTL   map[string]time.Time
	mu         sync.RWMutex
}

var (
	holidaysClient *client
	once           sync.Once
)

func getClient() *client {
	once.Do(func() {
		holidaysClient = &client{
			httpClient: &http.Client{Timeout: 10 * time.Second},
			cache:      make(map[string][]Holiday),
			cacheTTL:   make(map[string]time.Time),
		}
	})

	return holidaysClient
}
