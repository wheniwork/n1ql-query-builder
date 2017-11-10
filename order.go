package nqb

type direction string

const (
	asc  direction = " ASC"
	desc direction = " DESC"
)

func order(column string, dir direction) BuildFunc {
	return BuildFunc(func(buf *buffer) error {
		buf.WriteString(EscapeIdentifier(column))
		switch dir {
		case asc:
			buf.WriteString(string(asc))
		case desc:
			buf.WriteString(string(desc))
		}
		return nil
	})
}
