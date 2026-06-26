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

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tokensToAdd := tb.calculateTokensToAdd()

	if tokensToAdd > 0 {
		tb.lastRefill = time.Now()
	}

	tb.tokens += tokensToAdd
	if tb.tokens > tb.config.Capacity {
		tb.tokens = tb.config.Capacity
	}

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}

	return false
}

func (tb *TokenBucket) calculateTokensToAdd() int64 {
	elapsed := time.Since(tb.lastRefill)
	intervalsPassed := elapsed / tb.config.Interval
	tokensToAdd := int64(intervalsPassed) * tb.config.RefillRate
	return tokensToAdd
}

func (tb *TokenBucket) Tokens() int64 {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	return tb.tokens
}
