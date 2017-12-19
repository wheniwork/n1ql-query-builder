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

func (c *defaultSelectResult) Union() SelectClause {
	c.setElement(&unionElement{false, ""})
	return newDefaultSelectClause(c)
}

func (c *defaultSelectResult) UnionAll() SelectClause {
	c.setElement(&unionElement{true, ""})
	return newDefaultSelectClause(c)
}

func (c *defaultSelectResult) Intersect() SelectClause {
	c.setElement(&intersectElement{false, ""})
	return newDefaultSelectClause(c)
}

func (c *defaultSelectResult) IntersectAll() SelectClause {
	c.setElement(&intersectElement{true, ""})
	return newDefaultSelectClause(c)
}

func (c *defaultSelectResult) Except() SelectClause {
	c.setElement(&exceptElement{false, ""})
	return newDefaultSelectClause(c)
}

func (c *defaultSelectResult) ExceptAll() SelectClause {
	c.setElement(&exceptElement{true, ""})
	return newDefaultSelectClause(c)
}

func (c *defaultSelectResult) UnionClause(path SelectResult) SelectResult {
	c.setElement(&unionElement{false, path.String()})
	return newDefaultSelectResult(c)
}

func (c *defaultSelectResult) UnionAllClause(path SelectResult) SelectResult {
	c.setElement(&unionElement{true, path.String()})
	return newDefaultSelectResult(c)
}

func (c *defaultSelectResult) IntersectClause(path SelectResult) SelectResult {
	c.setElement(&intersectElement{false, path.String()})
	return newDefaultSelectResult(c)
}

func (c *defaultSelectResult) IntersectAllClause(path SelectResult) SelectResult {
	c.setElement(&intersectElement{true, path.String()})
	return newDefaultSelectResult(c)
}

func (c *defaultSelectResult) ExceptClause(path SelectResult) SelectResult {
	c.setElement(&exceptElement{false, path.String()})
	return newDefaultSelectResult(c)
}

func (c *defaultSelectResult) ExceptAllClause(path SelectResult) SelectResult {
	c.setElement(&exceptElement{true, path.String()})
	return newDefaultSelectResult(c)
}
