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
	for i := 0; i < 9; i++ {
		token, _ := bucket.NewToken(4)
		fmt.Printf("Generated Token %d: %s\n", i, token.Code)
	}

	// One more token that we will tamper with
	token, _ := bucket.NewToken(4)
	fmt.Printf("Generated Token 9: %s\n", token.Code)
	token.Code = "_" + token.Code[1:7] + "_"

	// Find the closest match for the faulty token
	match, distance := bucket.Resolve(token.Code)
	fmt.Printf("Best match for '%s' is token '%s' with distance %d\n", token.Code, match.Code, distance)
}
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

## Development

API docs can be found [here](http://godoc.org/github.com/muesli/toktok).

[![Build Status](https://secure.travis-ci.org/muesli/toktok.png)](http://travis-ci.org/muesli/toktok)
[![Coverage Status](https://coveralls.io/repos/github/muesli/toktok/badge.svg?branch=master)](https://coveralls.io/github/muesli/toktok?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/toktok)](http://goreportcard.com/report/muesli/toktok)
