package nqb

type AsPath interface {
	HintPath
	As(alias string) HintPath
}

type defaultAsPath struct {
	*defaultHintPath
}

func newDefaultAsPath(parent Path) *defaultAsPath {
	return &defaultAsPath{newDefaultHintPath(parent)}
}

func (p *defaultAsPath) As(alias string) HintPath {
	p.setElement(newAsAlement(alias))
	return newDefaultHintPath(p)
}

type asElement struct {
	as string
}

func newAsAlement(as string) *asElement {
	return &asElement{as}
}

func (e *asElement) Export() string {
	return "AS " + e.as
}
