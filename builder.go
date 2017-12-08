package nqb

import "bytes"

type BuildFunc func(buf *bytes.Buffer) error

func (b BuildFunc) Build(buf *bytes.Buffer) error {
	return b(buf)
}
