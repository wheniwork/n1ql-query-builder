package nqb

// element describes keywords in the N1QL DSL.
type element interface {
	export() string
}
