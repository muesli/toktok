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

var ()

func TestCodeGen(t *testing.T) {
	length := uint(8)
	bucket, err := NewBucket(length)
	if err != nil {
		t.Error("Error creating new token bucket:", err)
	}

	tok := bucket.NewToken(4)
	if len(tok.Code) != int(length) {
		t.Errorf("Wrong token length, expected %d, got %d", length, len(tok.Code))
	}
}
