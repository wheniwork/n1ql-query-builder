package nqb

import "bytes"

type direction string

const (
	asc  direction = " ASC"
	desc direction = " DESC"
)

func order(column string, dir direction) BuildFunc {
	return BuildFunc(func(buf *bytes.Buffer) error {
		buf.WriteString(escapeIdentifiers(column))
		switch dir {
		case asc:
			buf.WriteString(string(asc))
		case desc:
			buf.WriteString(string(desc))
		}
		return nil
	})
}
