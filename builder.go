package nqb

import "bytes"

// BuildFunc is a function that operates on a `bytes.Buffer`
type BuildFunc func(buf *bytes.Buffer) error

// Build executes a build function
func (b BuildFunc) Build(buf *bytes.Buffer) error {
	return b(buf)
}
