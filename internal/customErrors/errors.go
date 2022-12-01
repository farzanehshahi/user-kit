package customErrors

import (
	"errors"
)

var (
	ErrUserNotFound = errors.New("user doesn't exists")

	ErrUsernameAlreadyExists = errors.New("username already exists")

	ErrInvalidCredentials = errors.New("invalid email or password")
)
