package nqb

import "strconv"

type LimitPath interface {
	OffsetPath
	Limit(limit int) OffsetPath
}

type defaultLimitPath struct {
	*defaultOffsetPath
}

func newDefaultLimitPath(parent Path) *defaultLimitPath {
	return &defaultLimitPath{newDefaultOffsetPath(parent)}
}

func (p *defaultLimitPath) Limit(limit int) OffsetPath {
	p.setElement(newLimitElement(limit))
	return newDefaultOffsetPath(p)
}

type limitElement struct {
	limit int
}

func newLimitElement(limit int) *limitElement {
	return &limitElement{limit}
}

func (e *limitElement) Export() string {
	return "LIMIT " + strconv.Itoa(e.limit)
}
