package tPassword

import "errors"

var (
	ErrEmptyCharset  = errors.New("no charset selected")
	ErrInvalidLength = errors.New("invalid length")
)
