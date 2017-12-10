package nqb

type SelectResultPath interface {
	OrderByPath

	Union() SelectPath

	UnionAll() SelectPath

	Intersect() SelectPath

	IntersectAll() SelectPath

	Except() SelectPath

	ExceptAll() SelectPath

	UnionPath(path SelectResultPath) SelectResultPath

	UnionAllPath(path SelectResultPath) SelectResultPath

	IntersectPath(path SelectResultPath) SelectResultPath

	IntersectAllPath(path SelectResultPath) SelectResultPath

	ExceptPath(path SelectResultPath) SelectResultPath

	ExceptAllPath(path SelectResultPath) SelectResultPath
}

type defaultSelectResultPath struct {
	*defaultOrderByPath
}

func newDefaultSelectResultPath(parent Path) *defaultSelectResultPath {
	return &defaultSelectResultPath{newDefaultOrderByPath(parent)}
}

func (p *defaultSelectResultPath) Union() SelectPath {
	p.setElement(newUnionElement(false))
	return newDefaultSelectPath(p)
}

func (p *defaultSelectResultPath) UnionAll() SelectPath {
	p.setElement(newUnionElement(true))
	return newDefaultSelectPath(p)
}

func (p *defaultSelectResultPath) Intersect() SelectPath {
	p.setElement(newIntersectElement(false))
	return newDefaultSelectPath(p)
}

func (p *defaultSelectResultPath) IntersectAll() SelectPath {
	p.setElement(newIntersectElement(true))
	return newDefaultSelectPath(p)
}

func (p *defaultSelectResultPath) Except() SelectPath {
	p.setElement(newExceptElement(false))
	return newDefaultSelectPath(p)
}

func (p *defaultSelectResultPath) ExceptAll() SelectPath {
	p.setElement(newExceptElement(true))
	return newDefaultSelectPath(p)
}

func (p *defaultSelectResultPath) UnionPath(path SelectResultPath) SelectResultPath {
	p.setElement(newUnionElementPathStmt(false, path))
	return newDefaultSelectResultPath(p)
}

func (p *defaultSelectResultPath) UnionAllPath(path SelectResultPath) SelectResultPath {
	p.setElement(newUnionElementPathStmt(true, path))
	return newDefaultSelectResultPath(p)
}

func (p *defaultSelectResultPath) IntersectPath(path SelectResultPath) SelectResultPath {
	p.setElement(newIntersectElementPathStmt(false, path))
	return newDefaultSelectResultPath(p)
}

func (p *defaultSelectResultPath) IntersectAllPath(path SelectResultPath) SelectResultPath {
	p.setElement(newIntersectElementPathStmt(true, path))
	return newDefaultSelectResultPath(p)
}

func (p *defaultSelectResultPath) ExceptPath(path SelectResultPath) SelectResultPath {
	p.setElement(newExceptElementPathStmt(false, path))
	return newDefaultSelectResultPath(p)
}

func (p *defaultSelectResultPath) ExceptAllPath(path SelectResultPath) SelectResultPath {
	p.setElement(newExceptElementPathStmt(true, path))
	return newDefaultSelectResultPath(p)
}
