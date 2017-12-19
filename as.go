package nqb

type AsKeyword interface {
	UseIndexClause
	As(alias string) UseIndexClause
}

type defaultAsKeyword struct {
	*defaultUseIndexClause
}

func newDefaultAsKeyword(parent Statement) *defaultAsKeyword {
	return &defaultAsKeyword{newDefaultHintClause(parent)}
}

func (p *defaultAsKeyword) As(alias string) UseIndexClause {
	p.setElement(&asElement{alias})
	return newDefaultHintClause(p)
}

type asElement struct {
	as string
}

func (e *asElement) export() string {
	return "AS " + e.as
}
