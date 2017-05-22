/*
 * Human-friendly token generator
 *     Copyright (c) 2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

package toktok

import (
	"math/rand"
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
}

func NewBucket(tokenLength uint) (Bucket, error) {
	if tokenLength < 2 {
		return Bucket{}, ErrTokenLengthTooSmall
	}

	return Bucket{
		length: tokenLength,
		runes:  []rune("ABCDEFGHKLMNRSTWXZ234589"),
		Tokens: make(map[string]Token),
	}, nil
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

func (bucket *Bucket) NewToken(distance int) Token {
	if distance < 1 {
		return Token{}
	}

	var c string
	for {
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
	}

	token := Token{
		Code: c,
	}
	bucket.Tokens[token.Code] = token

	return token
}

func (bucket *Bucket) Resolve(code string) (Token, int) {
	var t Token
	distance := 65536

	for _, token := range bucket.Tokens {
		if hd := smetrics.WagnerFischer(code, token.Code, 1, 1, 2); hd <= distance {
			if hd == distance {
				// duplicate distance, ignore the previous result as it's not unique enough
				t = Token{}
			} else {
				t = token
				distance = hd

				if distance == 0 {
					// perfect match
					break
				}
			}
		}
	}

	return t, distance
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
