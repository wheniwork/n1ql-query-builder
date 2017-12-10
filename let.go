package nqb

import "bytes"

type Alias struct {
	alias    string
	original *Expression
}

func NewAliasExpr(alias string, original *Expression) *Alias {
	return &Alias{alias, original}
}

func NewAlias(alias, original string) *Alias {
	return NewAliasExpr(alias, X(original))
}

func (a *Alias) String() string {
	return a.alias + " = " + a.original.String()
}

type LetPath interface {
	WherePath

	Let(aliases ...*Alias) WherePath

	Join(from interface{}) JoinPath

	InnerJoin(from interface{}) JoinPath

	LeftJoin(from interface{}) JoinPath

	LeftOuterJoin(from interface{}) JoinPath

	Nest(from interface{}) NestPath

	InnerNest(from interface{}) NestPath

	LeftNest(from interface{}) NestPath

	LeftOuterNest(from interface{}) NestPath

	Unnest(from interface{}) UnnestPath

	InnerUnnest(from interface{}) UnnestPath

	LeftUnnest(from interface{}) UnnestPath

	LeftOuterUnnest(from interface{}) UnnestPath
}

type defaultLetPath struct {
	*defaultWherePath
}

func newDefaultLetPath(parent Path) *defaultLetPath {
	return &defaultLetPath{newDefaultWherePath(parent)}
}

func (p *defaultLetPath) Let(aliases ...*Alias) WherePath {
	p.setElement(&letElement{aliases})
	return newDefaultWherePath(p)
}

func (p *defaultLetPath) Join(from interface{}) JoinPath {
	p.setElement(&joinElement{defaultJoin, toString(from)})
	return newDefaultJoinPath(p)
}

func (p *defaultLetPath) InnerJoin(from interface{}) JoinPath {
	p.setElement(&joinElement{inner, toString(from)})
	return newDefaultJoinPath(p)
}

func (p *defaultLetPath) LeftJoin(from interface{}) JoinPath {
	p.setElement(&joinElement{left, toString(from)})
	return newDefaultJoinPath(p)
}

func (p *defaultLetPath) LeftOuterJoin(from interface{}) JoinPath {
	p.setElement(&joinElement{leftOuter, toString(from)})
	return newDefaultJoinPath(p)
}

func (p *defaultLetPath) Nest(from interface{}) NestPath {
	p.setElement(&nestElement{defaultJoin, toString(from)})
	return newDefaultNestPath(p)
}

func (p *defaultLetPath) InnerNest(from interface{}) NestPath {
	p.setElement(&nestElement{inner, toString(from)})
	return newDefaultNestPath(p)
}

func (p *defaultLetPath) LeftNest(from interface{}) NestPath {
	p.setElement(&nestElement{left, toString(from)})
	return newDefaultNestPath(p)
}

func (p *defaultLetPath) LeftOuterNest(from interface{}) NestPath {
	p.setElement(&nestElement{leftOuter, toString(from)})
	return newDefaultNestPath(p)
}

func (p *defaultLetPath) Unnest(from interface{}) UnnestPath {
	p.setElement(newUnnestElement(defaultJoin, toString(from)))
	return newDefaultUnnestPath(p)
}

func (p *defaultLetPath) InnerUnnest(from interface{}) UnnestPath {
	p.setElement(newUnnestElement(inner, toString(from)))
	return newDefaultUnnestPath(p)
}

func (p *defaultLetPath) LeftUnnest(from interface{}) UnnestPath {
	p.setElement(newUnnestElement(left, toString(from)))
	return newDefaultUnnestPath(p)
}

func (p *defaultLetPath) LeftOuterUnnest(from interface{}) UnnestPath {
	p.setElement(newUnnestElement(leftOuter, toString(from)))
	return newDefaultUnnestPath(p)
}

type letElement struct {
	aliases []*Alias
}

func (e *letElement) export() string {
	buf := bytes.NewBufferString("LET ")
	for i, alias := range e.aliases {
		buf.WriteString(alias.String())

		// todo improve?
		if i < len(e.aliases)-1 {
			buf.WriteString(", ")
		}
	}

	return buf.String()
}
