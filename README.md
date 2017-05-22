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
	for i := 0; i < 10; i++ {
		fmt.Printf("Generated Token %d: %s\n", i, bucket.NewToken(4).Code)
	}
}
```

## Result
```
Generated Token 0: DCN9TEX5
Generated Token 1: MHBMSXZK
Generated Token 2: D5HXWX2L
...
```

## Development

API docs can be found [here](http://godoc.org/github.com/muesli/toktok).

[![Build Status](https://secure.travis-ci.org/muesli/toktok.png)](http://travis-ci.org/muesli/toktok)
[![Coverage Status](https://coveralls.io/repos/github/muesli/toktok/badge.svg?branch=master)](https://coveralls.io/github/muesli/toktok?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/toktok)](http://goreportcard.com/report/muesli/toktok)
