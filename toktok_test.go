/*
 * Human-friendly token generator
 *     Copyright (c) 2017-2022, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

package toktok

import (
	"strings"
	"testing"
)

func TestTokenGen(t *testing.T) {
	length := uint(8)
	bucket, err := NewBucket(length)
	if err != nil {
		t.Error("Error creating new token bucket:", err)
	}

	tok, err := bucket.NewToken(4)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if len(tok) != int(length) {
		t.Errorf("Wrong token length, expected %d, got %d", length, len(tok))
	}
}

func TestTokenCount(t *testing.T) {
	bucket, _ := NewBucket(8)
	bucket.NewToken(1)

	if bucket.Count() != 1 {
		t.Errorf("Expected Count() to return 1, got %d", bucket.Count())
	}
}

func TestTokenLoad(t *testing.T) {
	code1, code2 := "ABCDEFGH", "IJKLMNOP"
	tokens := []string{code1, code2}
	bucket, _ := NewBucket(8)
	bucket.LoadTokens(tokens)

	if bucket.Count() != 2 {
		t.Errorf("Expected Count() to return 2, got %d", bucket.Count())
	}
	tok, _ := bucket.Resolve(code1)
	if tok != code1 {
		t.Errorf("Expected Token '%s', got '%s'", code1, tok)
	}
}

func TestTokenEstimations(t *testing.T) {
	bucket, _ := NewBucket(7)
	if bucket.EstimatedFillPercentage() != 0 {
		t.Errorf("Expected zero fill-rate estimate, got %f", bucket.EstimatedFillPercentage())
	}

	for i := 0; i < 5001; i++ {
		bucket.NewToken(4)
	}

	if bucket.EstimatedFillPercentage() <= 0 {
		t.Errorf("Expected positive fill-rate estimate, got %f", bucket.EstimatedFillPercentage())
	}
	if bucket.EstimatedTokenSpace() <= 0 {
		t.Errorf("Expected positive token space estimate, got %d", bucket.EstimatedTokenSpace())
	}
}

func TestTokenError(t *testing.T) {
	_, err := NewBucket(1)
	if err != ErrTokenLengthTooSmall {
		t.Errorf("Expected error %v, got %v", ErrTokenLengthTooSmall, err)
	}

	_, err = NewBucketWithRunes(8, "ABC")
	if err != ErrTooFewRunes {
		t.Errorf("Expected error %v, got %v", ErrTooFewRunes, err)
	}

	_, err = NewBucketWithRunes(8, "ABCDabcd")
	if err != ErrDupeRunes {
		t.Errorf("Expected error %v, got %v", ErrDupeRunes, err)
	}

	bucket, _ := NewBucket(4)
	_, err = bucket.NewToken(0)
	if err != ErrDistanceTooSmall {
		t.Errorf("Expected error %v, got %v", ErrDistanceTooSmall, err)
	}

	for i := 0; i < 256; i++ {
		_, err = bucket.NewToken(4)
		if err != nil {
			break
		}
	}
	if err != ErrTokenSpaceExhausted {
		t.Errorf("Expected error %v, got %v", ErrTokenSpaceExhausted, err)
	}
}

func TestTokenResolve(t *testing.T) {
	bucket, _ := NewBucket(8)

	var tok string
	for i := 0; i < 32; i++ {
		gtok, err := bucket.NewToken(4)
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}
		if i == 0 {
			tok = gtok
		}
	}

	ntok, dist := bucket.Resolve(tok)
	if ntok != tok {
		t.Errorf("Token mismatch, expected %v, got %v", tok, ntok)
	}
	if dist != 0 {
		t.Errorf("Wrong distance returned, expected 0, got %d", dist)
	}

	ntok, dist = bucket.Resolve(strings.ToLower(tok))
	if ntok != tok {
		t.Errorf("Lowercase token mismatch, expected %v, got %v", tok, ntok)
	}
	if dist != 0 {
		t.Errorf("Wrong distance returned, expected 0, got %d", dist)
	}
}

func TestTokenFaultyResolve(t *testing.T) {
	bucket, _ := NewBucket(8)

	var tok, ttok string
	for i := 0; i < 32; i++ {
		gtok, err := bucket.NewToken(4)
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}
		if i == 0 {
			tok = gtok
			ttok = tok
		}
	}

	// replace char in token
	ttok = " " + ttok[1:]

	ntok, dist := bucket.Resolve(ttok)
	if ntok != tok {
		t.Errorf("Token mismatch, expected %v, got %v", tok, ntok)
	}
	if dist != 2 {
		t.Errorf("Wrong distance returned, expected 2, got %d", dist)
	}

	// insert char in token
	ttok = tok + " "

	ntok, dist = bucket.Resolve(ttok)
	if ntok != tok {
		t.Errorf("Token mismatch, expected %v, got %v", tok, ntok)
	}
	if dist != 1 {
		t.Errorf("Wrong distance returned, expected 1, got %d", dist)
	}

	// remove char in token
	ttok = tok[1:]

	ntok, dist = bucket.Resolve(ttok)
	if ntok != tok {
		t.Errorf("Token mismatch, expected %v, got %v", tok, ntok)
	}
	if dist != 1 {
		t.Errorf("Wrong distance returned, expected 1, got %d", dist)
	}
}

func BenchmarkCodeGen(b *testing.B) {
	bucket, _ := NewBucket(8)

	for i := 0; i < b.N; i++ {
		bucket.NewToken(4)
	}
}
