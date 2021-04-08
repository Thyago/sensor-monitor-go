package util

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrAlreadyExists = errors.New("Already exists")
	ErrNotFound      = gorm.ErrRecordNotFound
	ErrBadRequest    = errors.New("Bad Request")
	ErrUnexpected    = errors.New("Unexpected error")

	ErrSelfLoop = errors.New("Self loop not allowed")
	ErrInvalid  = errors.New("Invalid")
)
