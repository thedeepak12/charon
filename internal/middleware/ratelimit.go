package middleware

import (
	"log"
	"net/http"
	"strconv"

	"github.com/thedeepak12/charon/internal/ratelimiter"
)

type RateLimitMiddleware struct {
	storage      *ratelimiter.Storage
	keyExtractor func(*http.Request) string
}

func NewRateLimitMiddleware(storage *ratelimiter.Storage, keyExtractor func(*http.Request) string) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		storage:      storage,
		keyExtractor: keyExtractor,
	}
}

func (m *RateLimitMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := m.keyExtractor(r)
		limiter := m.storage.GetOrCreate(key)

		allowed := limiter.Allow()

		w.Header().Set("X-RateLimit-Limit", strconv.FormatInt(limiter.Capacity(), 10))
		w.Header().Set("X-RateLimit-Remaining", strconv.FormatInt(limiter.Tokens(), 10))

		status := "ALLOWED"
		if !allowed {
			status = "DENIED"
		}
		log.Printf("[%s] %s %s | Tokens: %d/%d | Key: %s", status, r.Method, r.URL.Path, limiter.Tokens(), limiter.Capacity(), key)

		if !allowed {
			w.Header().Set("Retry-After", "1")
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ExtractIP(r *http.Request) string {
	var ip string

	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ip = xff
	} else if xri := r.Header.Get("X-Real-IP"); xri != "" {
		ip = xri
	} else {
		ip = r.RemoteAddr
	}

	for i := len(ip) - 1; i >= 0; i-- {
		if ip[i] == ':' {
			return ip[:i]
		}
	}

	return ip
}
