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
	for i := 0; i < 9; i++ {
		token, _ := bucket.NewToken(4)
		fmt.Printf("Generated Token %d: %s\n", i, token.Code)
	}

	// One more token that we will tamper with
	token, _ := bucket.NewToken(4)
	fmt.Printf("Generated Token 9: %s\n", token.Code)
	token.Code = "_" + token.Code[1:7] + "_"

	// Find the closest match for the faulty token
	match, distance := bucket.Resolve(token.Code)
	fmt.Printf("Best match for '%s' is token '%s' with distance %d\n", token.Code, match.Code, distance)
}
