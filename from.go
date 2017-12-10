package nqb

type FromPath interface {
	LetPath

	From(from string) AsPath

	FromExpr(expression *Expression) AsPath
}

type defaultFromPath struct {
	*defaultLetPath
}

func newDefaultFromPath(parent Path) *defaultFromPath {
	return &defaultFromPath{newDefaultLetPath(parent)}
}

func (p *defaultFromPath) From(from string) AsPath {
	p.setElement(&fromElement{from})
	return newDefaultAsPath(p)
}

func (p *defaultFromPath) FromExpr(from *Expression) AsPath {
	p.setElement(&fromElement{from.String()})
	return newDefaultAsPath(p)
}

type fromElement struct {
	from string
}

func (e *fromElement) export() string {
	return "FROM " + e.from
}
