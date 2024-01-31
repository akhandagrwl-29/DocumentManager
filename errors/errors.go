package errors

import "errors"

var (
	ErrNoInput                = errors.New("no Input provided")
	ErrInvalidCommand         = errors.New("invalid command")
	ErrUserAlreadyExists      = errors.New("user already exists")
	ErrDocumentAlreadyExists  = errors.New("document already exist")
	ErrDocumentNotExists      = errors.New("document does not exist")
	ErrUserNotAuthenticated   = errors.New("invalid user or password")
	ErrUserNotExists          = errors.New("user not exists")
	ErrDocumentAlreadyDeleted = errors.New("document already deleted")
	ErrInvalidVersion         = errors.New("invalid version")
)
