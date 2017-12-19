package nqb

type AsKeyword interface {
	UseIndexClause
	As(alias string) UseIndexClause
}

type defaultAsKeyword struct {
	*defaultUseIndexClause
}

func newDefaultAsKeyword(parent Statement) *defaultAsKeyword {
	return &defaultAsKeyword{newDefaultUseIndexClause(parent)}
}

func (k *defaultAsKeyword) As(alias string) UseIndexClause {
	k.setElement(&asElement{alias})
	return newDefaultUseIndexClause(k)
}

type asElement struct {
	as string
}

func (e *asElement) export() string {
	return "AS " + e.as
}
