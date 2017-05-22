/*
 * Human-friendly token generator
 *     Copyright (c) 2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

package toktok

import (
	"sync"
	"testing"
)

func BenchmarkCodeGen(b *testing.B) {
	bucket, _ := NewBucket(8)

	var finish sync.WaitGroup

	fn := func() {
		for i := 0; i < b.N; i++ {
			bucket.NewToken(4)
		}
		finish.Done()
	}

	finish.Add(1)
	go fn()
	finish.Wait()
}
