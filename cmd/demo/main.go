package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/thedeepak12/charon/internal/middleware"
	"github.com/thedeepak12/charon/internal/ratelimiter"
)

func main() {
	config := ratelimiter.DefaultConfig()
	config.Capacity = 10
	config.RefillRate = 1
	config.Interval = 1 * time.Second

	storage, err := ratelimiter.NewStorage(config)
	if err != nil {
		fmt.Println("Error creating storage:", err)
		return
	}

	rateLimitMiddleware := middleware.NewRateLimitMiddleware(storage, middleware.ExtractIP)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := middleware.ExtractIP(r)
		limiter := storage.GetOrCreate(key)

		response := fmt.Sprintf(
			"Request from: %s\nTokens: %d/%d\nActive limiters: %d",
			key,
			limiter.Tokens(),
			limiter.Capacity(),
			storage.Size(),
		)
		w.Write([]byte(response))
	})

	handler := rateLimitMiddleware.Handler(mux)

	fmt.Println("Server starting on :4000")
	http.ListenAndServe(":4000", handler)
}
