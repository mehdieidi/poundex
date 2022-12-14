package soundex

import "errors"

var (
	ErrPersianWordTooSmall = errors.New("persian word too small")
	ErrInvalidPersianWord  = errors.New("invalid persian word")
)
