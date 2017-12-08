package nqb

import "bytes"

type Builder interface {
	Build() error
}

type BuildFunc func(buf *bytes.Buffer) error

func (b BuildFunc) Build(buf *bytes.Buffer) error {
	return b(buf)
}
