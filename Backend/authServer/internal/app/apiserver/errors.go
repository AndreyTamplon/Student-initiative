package apiserver

import "errors"

var (
	ErrIncorrectEmailOrPassword = errors.New("incorrect email or password")
	ErrNotAuthenticated         = errors.New("not authenticated")
	ErrAlreadyConfirmed         = errors.New("already confirmed")
	ErrInvalidCode              = errors.New("invalid confirmation code")
)
