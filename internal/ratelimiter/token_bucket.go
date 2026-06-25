package ratelimiter

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	config     *Config
	tokens     int64
	lastRefill time.Time
	mu         sync.Mutex
}

func NewTokenBucket(config *Config) (*TokenBucket, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &TokenBucket{
		config:     config,
		tokens:     config.Capacity,
		lastRefill: time.Now(),
	}, nil
}

func (tb *TokenBucket) String() string {
	return fmt.Sprintf("TokenBucket{Capacity: %d, Tokens: %d, RefillRate: %d, Interval: %v}",
		tb.config.Capacity, tb.tokens, tb.config.RefillRate, tb.config.Interval)
}
