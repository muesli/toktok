toktok
======

[![Latest Release](https://img.shields.io/github/release/muesli/toktok.svg)](https://github.com/muesli/toktok/releases)
[![Build Status](https://travis-ci.org/muesli/toktok.svg?branch=master)](https://travis-ci.org/muesli/toktok)
[![Coverage Status](https://coveralls.io/repos/github/muesli/toktok/badge.svg?branch=master)](https://coveralls.io/github/muesli/toktok?branch=master)
[![Go ReportCard](https://goreportcard.com/badge/muesli/toktok)](https://goreportcard.com/report/muesli/toktok)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/toktok)

A human-friendly token generator

Creates tokens which avoid characters that can be easily misinterpreted, like '1' and 'I' or '8' and 'B', as well as
repeated characters within the token. It also compares newly generated tokens to all previously generated ones and
guarantees a safety distance between the tokens, so they become resilient to typos or other human entry errors.

## Installation

Make sure you have a working Go environment (Go 1.11 or higher is required).
See the [install instructions](https://golang.org/doc/install.html).

To install toktok, simply run:

    go get github.com/muesli/toktok

Compiling toktok is easy, simply run:

    git clone https://github.com/muesli/toktok.git
    cd toktok
    go build && go test -v

## Example
```go
package main

import (
	"fmt"

	"github.com/muesli/toktok"
)

func main() {
	// Generate a new token bucket.
	// Each generated token will be 8 characters long.
	bucket, _ := toktok.NewBucket(8)

	// Generate a bunch of tokens with a safety distance of 4.
	// Distance is calculated by insertion cost (1), deletion cost (1) and
	// substitution cost (2).
	for i := 0; i < 9; i++ {
		token, _ := bucket.NewToken(4)
		fmt.Printf("Generated Token %d: %s\n", i, token)
	}

	// One more token that we will tamper with.
	token, _ := bucket.NewToken(4)
	fmt.Printf("Generated Token 9: %s\n", token)
	token = "_" + token[1:7] + "_"

	// Find the closest match for the faulty token.
	match, distance := bucket.Resolve(token)
	fmt.Printf("Best match for '%s' is token '%s' with distance %d\n", token, match, distance)
}
```

## Result
```
Generated Token 0: J3KPC9YF
Generated Token 1: PXTWDC9P
Generated Token 2: WNANK4FU
...
Generated Token 9: Y3NCDFWN
Best match for '_3NCDFW_' is token 'Y3NCDFWN' with distance 4
```

## Feedback

Got some feedback or suggestions? Please open an issue or drop me a note!

* [Twitter](https://twitter.com/mueslix)
* [The Fediverse](https://mastodon.social/@fribbledom)
