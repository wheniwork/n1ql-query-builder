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
	p.setElement(newFromElement(from))
	return newDefaultAsPath(p)
}

func (p *defaultFromPath) FromExpr(from *Expression) AsPath {
	p.setElement(newFromElement(from.String()))
	return newDefaultAsPath(p)
}

type fromElement struct {
	from string
}

func newFromElement(from string) *fromElement {
	return &fromElement{from}
}

func (e *fromElement) Export() string {
	return "FROM " + e.from
}
