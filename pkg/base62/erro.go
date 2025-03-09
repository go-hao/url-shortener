package base62

import "errors"

var (
	ErrInvalidBase62Value = errors.New("invalid base62 value")
	ErrNotInit            = errors.New("base62 is not initialized")
	ErrInvalidSeed62      = errors.New("invalid seed62 string")
)
