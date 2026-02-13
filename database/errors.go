package database

import "errors"

var (
	ErrFileAlreadyExists = errors.New("database file already exists")
	ErrFileDoesNotExist  = errors.New("database file does not exist")
)
