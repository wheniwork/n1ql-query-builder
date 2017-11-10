package nqb

type Builder interface {
	Build() error
}

type BuildFunc func(buf *buffer) error

func (b BuildFunc) Build(buf *buffer) error {
	return b(buf)
}
