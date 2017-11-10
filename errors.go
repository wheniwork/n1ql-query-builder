package nqb

import "errors"

var (
	ErrNotFound           = errors.New("nqb: not found")
	ErrNotSupported       = errors.New("nqb: not supported")
	ErrKeyspaceNotSpecified  = errors.New("nqb: keyspace not specified")
	ErrColumnNotSpecified = errors.New("nqb: columns not specified")
)
