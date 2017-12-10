package nqb

type FromPath interface {
	LetPath

	From(from interface{}) AsPath
}

type defaultFromPath struct {
	*defaultLetPath
}

func newDefaultFromPath(parent Path) *defaultFromPath {
	return &defaultFromPath{newDefaultLetPath(parent)}
}

func (p *defaultFromPath) From(from interface{}) AsPath {
	p.setElement(&fromElement{toString(from)})
	return newDefaultAsPath(p)
}

type fromElement struct {
	from string
}

func (e *fromElement) export() string {
	return "FROM " + e.from
}
