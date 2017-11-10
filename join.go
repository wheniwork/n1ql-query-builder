package nqb

type joinType uint8

const (
	inner joinType = iota
	left
	right
	full
)

func join(t joinType, keyspace interface{}, on interface{}) BuildFunc {
	return BuildFunc(func(buf *buffer) error {
		buf.WriteString(" ")

		switch t {
		case left:
			buf.WriteString("LEFT ")
		case right:
			buf.WriteString("RIGHT ")
		case full:
			buf.WriteString("FULL ")
		}

		buf.WriteString("JOIN ")

		switch keyspace := keyspace.(type) {
		case string:
			buf.WriteString(EscapeIdentifier(keyspace))
		default:
			buf.WriteString(placeholder)
			buf.WriteValue(keyspace)
		}

		buf.WriteString(" ON ")

		switch on := on.(type) {
		case string:
			buf.WriteString(on)
		case Builder:
			buf.WriteString(placeholder)
			buf.WriteValue(on)
		}

		return nil
	})
}
