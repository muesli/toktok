/*
 * Human-friendly token generator
 *     Copyright (c) 2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

package toktok

import (
	"testing"
)

func TestCodeGen(t *testing.T) {
	length := uint(8)
	bucket, err := NewBucket(length)
	if err != nil {
		t.Error("Error creating new token bucket:", err)
	}

	tok, err := bucket.NewToken(4)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if len(tok.Code) != int(length) {
		t.Errorf("Wrong token length, expected %d, got %d", length, len(tok.Code))
	}
	if bucket.Count() != 1 {
		t.Errorf("Expected Count() to return 1, got %d", bucket.Count())
	}
}

func TestCodeError(t *testing.T) {
	_, err := NewBucket(1)
	if err != ErrTokenLengthTooSmall {
		t.Errorf("Expected error %v, got %v", ErrTokenLengthTooSmall, err)
	}

	_, err = NewBucketWithRunes(8, "foo")
	if err != ErrTooFewRunes {
		t.Errorf("Expected error %v, got %v", ErrTooFewRunes, err)
	}
}

func TestCodeResolve(t *testing.T) {
	bucket, err := NewBucket(8)
	if err != nil {
		t.Error("Error creating new token bucket:", err)
	}

	var tok Token
	for i := 0; i < 32; i++ {
		gtok, err := bucket.NewToken(4)
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}
		if i == 0 {
			tok = gtok
		}
	}

	ntok, dist := bucket.Resolve(tok.Code)
	if ntok != tok {
		t.Errorf("Token mismatch, expected %v, got %v", tok, ntok)
	}
	if dist != 0 {
		t.Errorf("Wrong distance returned, expected 0, got %d", dist)
	}
}

func TestCodeFaultyResolve(t *testing.T) {
	bucket, err := NewBucket(8)
	if err != nil {
		t.Error("Error creating new token bucket:", err)
	}

	var tok Token
	var ttok Token
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
	ttok.Code = " " + ttok.Code[1:]

	ntok, dist := bucket.Resolve(ttok.Code)
	if ntok != tok {
		t.Errorf("Token mismatch, expected %v, got %v", tok, ntok)
	}
	if dist != 2 {
		t.Errorf("Wrong distance returned, expected 2, got %d", dist)
	}

	// insert char in token
	ttok.Code = tok.Code + " "

	ntok, dist = bucket.Resolve(ttok.Code)
	if ntok != tok {
		t.Errorf("Token mismatch, expected %v, got %v", tok, ntok)
	}
	if dist != 1 {
		t.Errorf("Wrong distance returned, expected 1, got %d", dist)
	}

	// remove char in token
	ttok.Code = tok.Code[1:]

	ntok, dist = bucket.Resolve(ttok.Code)
	if ntok != tok {
		t.Errorf("Token mismatch, expected %v, got %v", tok, ntok)
	}
	if dist != 1 {
		t.Errorf("Wrong distance returned, expected 1, got %d", dist)
	}
}
