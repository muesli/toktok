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

type Token struct {
	Code string
}

type Bucket struct {
	length uint
	runes  []rune
	Tokens map[string]Token

	tries uint64

	sync.RWMutex
}

func NewBucket(tokenLength uint) (Bucket, error) {
	return NewBucketWithRunes(tokenLength, "ACDEFHJKLMNPRSTUWXY3469")
}

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
		Tokens: make(map[string]Token),
	}, nil
}

func (bucket *Bucket) NewToken(distance int) (Token, error) {
	if distance < 1 {
		return Token{}, ErrDistanceTooSmall
	}
	if bucket.EstimatedFillPercentage() > 97.0 {
		return Token{}, ErrTokenSpaceExhausted
	}

	bucket.Lock()
	defer bucket.Unlock()

	var c string
	i := 0
	for {
		i++
		c = GenerateToken(bucket.length, bucket.runes)

		dupe := false
		for _, token := range bucket.Tokens {
			if hd := smetrics.WagnerFischer(c, token.Code, 1, 1, 2); hd <= distance {
				dupe = true
				break
			}
		}
		if !dupe {
			break
		}
		if i > 1000 {
			return Token{}, ErrTokenSpaceExhausted
		}
	}

	token := Token{
		Code: c,
	}
	bucket.Tokens[token.Code] = token
	bucket.tries += uint64(i)

	return token, nil
}

func (bucket *Bucket) Resolve(code string) (Token, int) {
	distance := 65536

	bucket.RLock()
	defer bucket.RUnlock()

	// try to find a perfect match first
	t, ok := bucket.Tokens[code]
	if ok {
		return t, 0
	}

	// find the closest match
	for _, token := range bucket.Tokens {
		if hd := smetrics.WagnerFischer(code, token.Code, 1, 1, 2); hd <= distance {
			if hd == distance {
				// duplicate distance, ignore the previous result as it's not unique enough
				t = Token{}
			} else {
				t = token
				distance = hd
			}
		}
	}

	return t, distance
}

func (bucket *Bucket) EstimatedFillPercentage() float64 {
	bucket.Lock()
	defer bucket.Unlock()

	return 100.0 - (100.0 / (float64(bucket.tries) / float64(len(bucket.Tokens))))
}

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
