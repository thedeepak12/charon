package main

import (
	"fmt"
	"time"

	"github.com/thedeepak12/charon/internal/ratelimiter"
)

func main() {
	config := ratelimiter.DefaultConfig()
	config.Capacity = 20
	config.RefillRate = 1
	config.Interval = 1 * time.Second

	bucket, err := ratelimiter.NewLimiter(config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Bucket initialized:\n%v\n", bucket)

	for i := 0; i < 50; i++ {
		allowed := bucket.Allow()
		fmt.Printf("Request %d: Allowed = %v\n Tokens = %d\n", i+1, allowed, bucket.Tokens())
		time.Sleep(500 * time.Millisecond)
	}
}
