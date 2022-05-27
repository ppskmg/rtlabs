package store

import "errors"

var (
	// ErrRecordNotFound ...
	ErrRecordNotFound = errors.New("record not found")
	ErrNoRows         = "sql: no rows in result set"
)
