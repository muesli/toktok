toktok
======

Human-friendly token generator.

Creates tokens which avoid characters that can be easily misinterpreted, like '1' and 'I' or '8' and 'B'.
It also compares newly generated tokens to all previously generated ones and guarantees a safety distance
between the tokens, so they become resilient to typos or other human entry errors.

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
	// Genrate a new token bucket. Each generated token will be 8 characters long
	bucket, _ := toktok.NewBucket(8)

	// Generate a bunch of tokens with a safety distance of 4
	for i := 0; i < 10; i++ {
		fmt.Println("Generated Token:", bucket.NewToken(4).Code)
	}
}
```

## Development

API docs can be found [here](http://godoc.org/github.com/muesli/toktok).

[![Build Status](https://secure.travis-ci.org/muesli/toktok.png)](http://travis-ci.org/muesli/toktok)
[![Coverage Status](https://coveralls.io/repos/github/muesli/toktok/badge.svg?branch=master)](https://coveralls.io/github/muesli/toktok?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/toktok)](http://goreportcard.com/report/muesli/toktok)
