package main

import (
	"fmt"

	"github.com/muesli/toktok"
)

func main() {
	for n := uint(4); n < 10; n++ {
		fmt.Println("Probing token space for length", n)
		fmt.Println("================================")
		for dist := 3; dist <= 5; dist++ {
			fmt.Printf("Probing with distance %d\n", dist)

			// Generate a new token bucket. Each generated token will be n characters long
			bucket, _ := toktok.NewBucket(n)

			// Generate a bunch of tokens with a safety distance of 4
			// Distance is calculated by insertion cost (1), deletion cost (1) and substitution cost (2)
			for i := 0; bucket.EstimatedFillPercentage() < 2.00 || i < 20000; i++ {
				_, err := bucket.NewToken(dist)
				if err != nil {
					break
				}

				if i >= 1024 && (i&(i-1)) == 0 {
					fmt.Printf("Generated %d tokens, estimated space for %d tokens (%.2f%% full)\n",
						i, bucket.EstimatedTokenSpace(), bucket.EstimatedFillPercentage())
				}
			}

			fmt.Printf("Finished probing for length %d with distance %d: estimated space for %d tokens\n\n", n, dist, bucket.EstimatedTokenSpace())
		}
		fmt.Println()
	}
}
