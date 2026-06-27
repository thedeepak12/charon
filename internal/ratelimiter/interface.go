package ratelimiter

type Limiter interface {
	Allow() bool
	Tokens() int64
	Capacity() int64
	String() string
}

func NewLimiter(config *Config) (Limiter, error) {
	return NewTokenBucket(config)
}
