package go_morse

import (
	"errors"
)

var (
	ErrNilMorseCodeHandler  = errors.New("morse code handler is nil")
	ErrUnicodeAlreadyExists = "unicode already exists: %c, %U"
	ErrUnicodeDoesNotExist  = "unicode does not exist: %c, %U"
	ErrCodeAlreadyExists    = "code already exists: %s"
	ErrCodeDoesNotExist     = "code does not exist: %s"
)
