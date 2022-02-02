/*
 * Human-friendly token generator
 *     Copyright (c) 2017-2022, Christian Muehlhaeuser <muesli@gmail.com>
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
	// ErrDupeRunes gets returned when the set of runes contains a dupe
	ErrDupeRunes = errors.New("Dupe in runes")
	// ErrDistanceTooSmall gets returned when the required distance is too small
	ErrDistanceTooSmall = errors.New("Distance must be at least 1")
	// ErrTokenSpaceExhausted gets returned when the token space has been exhausted
	ErrTokenSpaceExhausted = errors.New("Token space exhausted. Use longer tokens, more runes or a smaller distance")
)
