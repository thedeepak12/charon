package ratelimiter

import "sync"

type Storage struct {
	limiters map[string]Limiter
	config   *Config
	mu       sync.RWMutex
}

func NewStorage(config *Config) (*Storage, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &Storage{
		limiters: make(map[string]Limiter),
		config:   config,
	}, nil
}

func (s *Storage) GetOrCreate(key string) Limiter {
	s.mu.Lock()
	defer s.mu.Unlock()

	if limiter, exists := s.limiters[key]; exists {
		return limiter
	}

	limiter, _ := NewLimiter(s.config)
	s.limiters[key] = limiter
	return limiter
}

func (s *Storage) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.limiters)
}
