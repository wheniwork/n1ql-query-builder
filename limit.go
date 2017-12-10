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
	p.setElement(&limitElement{limit})
	return newDefaultOffsetPath(p)
}

type limitElement struct {
	limit int
}

func (e *limitElement) export() string {
	return "LIMIT " + strconv.Itoa(e.limit)
}
