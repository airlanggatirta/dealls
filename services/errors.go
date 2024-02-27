// services/errors.go

package services

import "errors"

var (
	ErrUsernameExists     = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credential")
)
