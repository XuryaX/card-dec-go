package exceptions

import "errors"

var (
	ErrDeckNotFound      = errors.New("deck not found")
	ErrInsufficientCards = errors.New("insufficient cards in deck")
	InvalidCount         = errors.New("InvalidCount")
)
