package main

import (
	"fmt"

	"github.com/thedeepak12/charon/internal/ratelimiter"
)

func main() {
	config := ratelimiter.DefaultConfig()

	bucket, err := ratelimiter.NewTokenBucket(config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Bucket initialized:\n%v\n", bucket)
}
