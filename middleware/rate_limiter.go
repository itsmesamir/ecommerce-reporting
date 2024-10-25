package middleware

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu          sync.Mutex
	requests    map[string]int
	lastRequest map[string]time.Time
	limit       int
	window      time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests:    make(map[string]int),
		lastRequest: make(map[string]time.Time),
		limit:       limit,
		window:      window,
	}
}

func (rl *RateLimiter) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ip := r.RemoteAddr // Get the client's IP address

	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Reset the request count if the time window has passed
	if lastRequestTime, exists := rl.lastRequest[ip]; exists && time.Since(lastRequestTime) > rl.window {
		rl.requests[ip] = 0
	}

	// Update the last request time
	rl.lastRequest[ip] = time.Now()

	rl.requests[ip]++

	// Check if the limit has been exceeded
	if rl.requests[ip] > rl.limit {
		http.Error(w, "Too many requests. Please try again later.", http.StatusTooManyRequests)
		return
	}

	next(w, r)
}
