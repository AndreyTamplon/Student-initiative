package store

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrAlreadySigned  = errors.New("already signed")
)
