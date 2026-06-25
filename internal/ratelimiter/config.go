package ratelimiter

import (
	"fmt"
	"time"
)

type Config struct {
	Capacity   int64
	RefillRate int64
	Interval   time.Duration
}

func (c *Config) Validate() error {
	if c.Capacity <= 0 {
		return fmt.Errorf("capacity must be positive")
	}

	if c.RefillRate <= 0 {
		return fmt.Errorf("refill rate must be positive")
	}

	if c.Interval <= 0 {
		return fmt.Errorf("interval must be positive")
	}

	return nil
}

func DefaultConfig() *Config {
	return &Config{
		Capacity:   100,
		RefillRate: 10,
		Interval:   time.Second,
	}
}
