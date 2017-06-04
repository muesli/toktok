/*
 * Human-friendly token generator
 *     Copyright (c) 2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

package toktok

import (
	"math/rand"
	"sync"
	"time"

	"github.com/xrash/smetrics"
)

// Bucket tracks all the generated tokens and lets you create new, unique tokens
type Bucket struct {
	length uint
	runes  []rune

	tokens map[string]bool
	tries  []uint64

	sync.RWMutex
}

// NewBucket returns a new bucket, which will contain tokens of tokenLength
func NewBucket(tokenLength uint) (Bucket, error) {
	return NewBucketWithRunes(tokenLength, "ACDEFHJKLMNPRSTUWXY3469")
}

// NewBucketWithRunes returns a new bucket and let's you define which runes will be used for token generation
func NewBucketWithRunes(tokenLength uint, runes string) (Bucket, error) {
	if tokenLength < 2 {
		return Bucket{}, ErrTokenLengthTooSmall
	}
	if len(runes) < 4 {
		return Bucket{}, ErrTooFewRunes
	}

	return Bucket{
		length: tokenLength,
		runes:  []rune(runes),
		tokens: make(map[string]bool),
	}, nil
}

// LoadTokens adds previously generated tokens to the Bucket
func (bucket *Bucket) LoadTokens(tokens []string) {
	bucket.Lock()
	defer bucket.Unlock()

	for _, v := range tokens {
		bucket.tokens[v] = true
	}
}

// NewToken returns a new token with a minimal safety distance to all other existing tokens
func (bucket *Bucket) NewToken(distance int) (string, error) {
	if distance < 1 {
		return "", ErrDistanceTooSmall
	}

	bucket.Lock()
	defer bucket.Unlock()

	var c string
	i := 0
	for {
		i++
		c = GenerateToken(bucket.length, bucket.runes)

		dupe := false
		for token := range bucket.tokens {
			if hd := smetrics.WagnerFischer(c, token, 1, 1, 2); hd <= distance {
				dupe = true
				break
			}
		}
		if !dupe {
			break
		}
		if i > 100 {
			return "", ErrTokenSpaceExhausted
		}
	}

	bucket.tokens[c] = true

	bucket.tries = append(bucket.tries, uint64(i))
	if len(bucket.tries) > 5000 {
		bucket.tries = bucket.tries[1:]
	}

	return c, nil
}

// Resolve tries to find the matching original token for a potentially corrupted token
func (bucket *Bucket) Resolve(code string) (string, int) {
	distance := 65536

	bucket.RLock()
	defer bucket.RUnlock()

	// try to find a perfect match first
	_, ok := bucket.tokens[code]
	if ok {
		return code, 0
	}

	var t string
	// find the closest match
	for token := range bucket.tokens {
		if hd := smetrics.WagnerFischer(code, token, 1, 1, 2); hd <= distance {
			if hd == distance {
				// duplicate distance, ignore the previous result as it's not unique enough
				t = ""
			} else {
				t = token
				distance = hd
			}
		}
	}

	return t, distance
}

// Count returns how many tokens are currently in this Bucket
func (bucket *Bucket) Count() uint64 {
	bucket.Lock()
	defer bucket.Unlock()

	return uint64(len(bucket.tokens))
}

// EstimatedFillPercentage returns how full the Bucket approximately is
func (bucket *Bucket) EstimatedFillPercentage() float64 {
	bucket.Lock()
	defer bucket.Unlock()

	if len(bucket.tries) == 0 {
		return 0
	}

	tries := uint64(0)
	for _, v := range bucket.tries {
		tries += v
	}

	return 100.0 - (100.0 / (float64(tries) / float64(len(bucket.tries))))
}

// EstimatedTokenSpace returns the total estimated token space available in this Bucket
func (bucket *Bucket) EstimatedTokenSpace() uint64 {
	return uint64(float64(bucket.Count()) * (100.0 / bucket.EstimatedFillPercentage()))
}

// GenerateToken generates a new token of length n with the defined rune-set letterRunes
func GenerateToken(n uint, letterRunes []rune) string {
	l := len(letterRunes)
	b := make([]rune, n)

	for i := range b {
		var lastrune rune
		if i > 0 {
			lastrune = b[i-1]
		}
		b[i] = lastrune
		for lastrune == b[i] {
			b[i] = letterRunes[rand.Intn(l)]
		}
	}

	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
