package nqb

const EscapeChar = "`"

type Element interface {
	Export() string
}
