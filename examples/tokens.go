package main

import (
	"fmt"

	"github.com/muesli/toktok"
)

func main() {
	// Generate a new token bucket. Each generated token will be 8 characters long
	bucket, _ := toktok.NewBucket(8)

	// Generate a bunch of tokens with a safety distance of 4
	// Distance is calculated by insertion cost (1), deletion cost (1) and substitution cost (2)
	for i := 0; i < 10; i++ {
		fmt.Printf("Generated Token %d: %s\n", i, bucket.NewToken(4).Code)
	}
}
