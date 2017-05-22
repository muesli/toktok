toktok
======

A human-friendly token generator

Creates tokens which avoid characters that can be easily misinterpreted, like '1' and 'I' or '8' and 'B', as well as
repeated characters within the token. It also compares newly generated tokens to all previously generated ones and
guarantees a safety distance between the tokens, so they become resilient to typos or other human entry errors.

## Installation

Make sure you have a working Go environment. See the [install instructions](http://golang.org/doc/install.html).

To install toktok, simply run:

    go get github.com/muesli/toktok

To compile it from source:

    cd $GOPATH/src/github.com/muesli/toktok
    go get -u -v
    go build && go test -v

## Example
```go
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
	for i := 1; i < 10; i++ {
		fmt.Printf("Generated Token %d: %s\n", i, bucket.NewToken(4).Code)
	}

	// One more token that we will tamper with
	code := bucket.NewToken(4).Code
	fmt.Printf("Generated Token 10: %s\n", code)
	code = "_" + code[1:]

	// Find the closest match for the faulty token
	token, distance := bucket.Resolve(code)
	fmt.Printf("Best match for '%s' is token '%s' with distance %d\n", code, token.Code, distance)
}
```

## Result
```
Generated Token 1: DCN9TEX5
Generated Token 2: MHBMSXZK
Generated Token 3: D5HXWX2L
...
Generated Token 10: C2B9ELX4
Closest match for '_2B9ELX4' is token 'C2B9ELX4' with distance 1
```

## Development

API docs can be found [here](http://godoc.org/github.com/muesli/toktok).

[![Build Status](https://secure.travis-ci.org/muesli/toktok.png)](http://travis-ci.org/muesli/toktok)
[![Coverage Status](https://coveralls.io/repos/github/muesli/toktok/badge.svg?branch=master)](https://coveralls.io/github/muesli/toktok?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/toktok)](http://goreportcard.com/report/muesli/toktok)
