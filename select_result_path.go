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
	p.setElement(&unionElement{false, ""})
	return newDefaultSelectPath(p)
}

func (p *defaultSelectResultPath) UnionAll() SelectPath {
	p.setElement(&unionElement{true, ""})
	return newDefaultSelectPath(p)
}

func (p *defaultSelectResultPath) Intersect() SelectPath {
	p.setElement(&intersectElement{false, ""})
	return newDefaultSelectPath(p)
}

func (p *defaultSelectResultPath) IntersectAll() SelectPath {
	p.setElement(&intersectElement{true, ""})
	return newDefaultSelectPath(p)
}

func (p *defaultSelectResultPath) Except() SelectPath {
	p.setElement(&exceptElement{false, ""})
	return newDefaultSelectPath(p)
}

func (p *defaultSelectResultPath) ExceptAll() SelectPath {
	p.setElement(&exceptElement{true, ""})
	return newDefaultSelectPath(p)
}

func (p *defaultSelectResultPath) UnionPath(path SelectResultPath) SelectResultPath {
	p.setElement(&unionElement{false, path.String()})
	return newDefaultSelectResultPath(p)
}

func (p *defaultSelectResultPath) UnionAllPath(path SelectResultPath) SelectResultPath {
	p.setElement(&unionElement{true, path.String()})
	return newDefaultSelectResultPath(p)
}

func (p *defaultSelectResultPath) IntersectPath(path SelectResultPath) SelectResultPath {
	p.setElement(&intersectElement{false, path.String()})
	return newDefaultSelectResultPath(p)
}

func (p *defaultSelectResultPath) IntersectAllPath(path SelectResultPath) SelectResultPath {
	p.setElement(&intersectElement{true, path.String()})
	return newDefaultSelectResultPath(p)
}

func (p *defaultSelectResultPath) ExceptPath(path SelectResultPath) SelectResultPath {
	p.setElement(&exceptElement{false, path.String()})
	return newDefaultSelectResultPath(p)
}

func (p *defaultSelectResultPath) ExceptAllPath(path SelectResultPath) SelectResultPath {
	p.setElement(&exceptElement{true, path.String()})
	return newDefaultSelectResultPath(p)
}
