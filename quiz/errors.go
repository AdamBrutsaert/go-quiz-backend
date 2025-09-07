package quiz

import "errors"

var (
	ErrInvalidCommand    = errors.New("invalid command")
	ErrInvalidName       = errors.New("invalid name")
	ErrNameAlreadyTaken  = errors.New("name already taken")
	ErrAlreadyRegistered = errors.New("already registered")
	ErrNotOwner          = errors.New("not the lobby owner")
)
