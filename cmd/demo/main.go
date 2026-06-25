package main

import (
	"fmt"
	
	"github.com/thedeepak12/charon/internal/ratelimiter"
)

func main() {
	bucket := ratelimiter.NewTokenBucket()
	
	fmt.Printf("Bucket initialized: %v\n", bucket)
}
