package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"finflow/internal/http_helpers"
)

type rateLimiter struct {
	visitors map[string]*visitor
	mu       sync.RWMutex
	rate     int           // requests per minute
	burst    int           // maximum burst size
	cleanup  time.Duration // cleanup interval
}

type visitor struct {
	tokens     int
	lastSeen   time.Time
	lastRefill time.Time
}

var limiter *rateLimiter

func init() {
	limiter = &rateLimiter{
		visitors: make(map[string]*visitor),
		rate:     60,          // 60 requests per minute
		burst:    10,          // allow bursts of 10 requests
		cleanup:  time.Minute, // cleanup every minute
	}

	// Start cleanup goroutine
	go limiter.cleanupVisitors()
}

func (rl *rateLimiter) getVisitor(ip string) *visitor {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	if !exists {
		v = &visitor{
			tokens:     rl.burst,
			lastSeen:   time.Now(),
			lastRefill: time.Now(),
		}
		rl.visitors[ip] = v
	}

	return v
}

func (rl *rateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[ip]
	if !exists {
		v = &visitor{
			tokens:     rl.burst - 1,
			lastSeen:   time.Now(),
			lastRefill: time.Now(),
		}
		rl.visitors[ip] = v
		return true
	}

	v.lastSeen = time.Now()

	// Refill tokens based on time passed
	now := time.Now()
	elapsed := now.Sub(v.lastRefill)
	tokensToAdd := int(elapsed.Minutes() * float64(rl.rate))

	if tokensToAdd > 0 {
		v.tokens += tokensToAdd
		if v.tokens > rl.burst {
			v.tokens = rl.burst
		}
		v.lastRefill = now
	}

	if v.tokens > 0 {
		v.tokens--
		return true
	}

	return false
}

func (rl *rateLimiter) cleanupVisitors() {
	ticker := time.NewTicker(rl.cleanup)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

// RateLimitMiddleware limits the number of requests from a single IP
func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}

		if !limiter.allow(ip) {
			http_helpers.WriteJSONError(w, http.StatusTooManyRequests, "rate limit exceeded")
			return
		}

		next.ServeHTTP(w, r)
	}
}
