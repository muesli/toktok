/*
 * Human-friendly token generator
 *     Copyright (c) 2017, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE
 */

package toktok

import (
	"errors"
)

var (
	// ErrTokenLengthTooSmall gets returned when the token length is too small
	ErrTokenLengthTooSmall = errors.New("Token length is too small")
	// ErrTooFewRunes gets returned when the set of runes is too small
	ErrTooFewRunes = errors.New("Not enough runes")
)
