package nqb

import "bytes"

type buffer struct {
	bytes.Buffer
	v []interface{}
}

func (b *buffer) WriteValue(v ...interface{}) error {
	b.v = append(b.v, v...)
	return nil
}

func (b *buffer) Value() []interface{} {
	return b.v
}
