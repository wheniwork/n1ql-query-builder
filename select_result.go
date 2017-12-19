package nqb

type SelectResult interface {
	OrderByClause

	Union() SelectClause

	UnionAll() SelectClause

	Intersect() SelectClause

	IntersectAll() SelectClause

	Except() SelectClause

	ExceptAll() SelectClause

	UnionClause(path SelectResult) SelectResult

	UnionAllClause(path SelectResult) SelectResult

	IntersectClause(path SelectResult) SelectResult

	IntersectAllClause(path SelectResult) SelectResult

	ExceptClause(path SelectResult) SelectResult

	ExceptAllClause(path SelectResult) SelectResult
}

type defaultSelectResult struct {
	*defaultOrderByClause
}

func newDefaultSelectResult(parent Statement) *defaultSelectResult {
	return &defaultSelectResult{newDefaultOrderByClause(parent)}
}

func (p *defaultSelectResult) Union() SelectClause {
	p.setElement(&unionElement{false, ""})
	return newDefaultSelectClause(p)
}

func (p *defaultSelectResult) UnionAll() SelectClause {
	p.setElement(&unionElement{true, ""})
	return newDefaultSelectClause(p)
}

func (p *defaultSelectResult) Intersect() SelectClause {
	p.setElement(&intersectElement{false, ""})
	return newDefaultSelectClause(p)
}

func (p *defaultSelectResult) IntersectAll() SelectClause {
	p.setElement(&intersectElement{true, ""})
	return newDefaultSelectClause(p)
}

func (p *defaultSelectResult) Except() SelectClause {
	p.setElement(&exceptElement{false, ""})
	return newDefaultSelectClause(p)
}

func (p *defaultSelectResult) ExceptAll() SelectClause {
	p.setElement(&exceptElement{true, ""})
	return newDefaultSelectClause(p)
}

func (p *defaultSelectResult) UnionClause(path SelectResult) SelectResult {
	p.setElement(&unionElement{false, path.String()})
	return newDefaultSelectResult(p)
}

func (p *defaultSelectResult) UnionAllClause(path SelectResult) SelectResult {
	p.setElement(&unionElement{true, path.String()})
	return newDefaultSelectResult(p)
}

func (p *defaultSelectResult) IntersectClause(path SelectResult) SelectResult {
	p.setElement(&intersectElement{false, path.String()})
	return newDefaultSelectResult(p)
}

func (p *defaultSelectResult) IntersectAllClause(path SelectResult) SelectResult {
	p.setElement(&intersectElement{true, path.String()})
	return newDefaultSelectResult(p)
}

func (p *defaultSelectResult) ExceptClause(path SelectResult) SelectResult {
	p.setElement(&exceptElement{false, path.String()})
	return newDefaultSelectResult(p)
}

func (p *defaultSelectResult) ExceptAllClause(path SelectResult) SelectResult {
	p.setElement(&exceptElement{true, path.String()})
	return newDefaultSelectResult(p)
}
