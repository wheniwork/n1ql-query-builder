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

type LetClause interface {
	WhereClause

	Let(aliases ...*Alias) WhereClause

	Join(from interface{}) JoinClause

	InnerJoin(from interface{}) JoinClause

	LeftJoin(from interface{}) JoinClause

	LeftOuterJoin(from interface{}) JoinClause

	Nest(from interface{}) NestClause

	InnerNest(from interface{}) NestClause

	LeftNest(from interface{}) NestClause

	LeftOuterNest(from interface{}) NestClause

	Unnest(from interface{}) UnnestClause

	InnerUnnest(from interface{}) UnnestClause

	LeftUnnest(from interface{}) UnnestClause

	LeftOuterUnnest(from interface{}) UnnestClause
}

type defaultLetClause struct {
	*defaultWhereClause
}

func newDefaultLetClause(parent Statement) *defaultLetClause {
	return &defaultLetClause{newDefaultWhereClause(parent)}
}

func (p *defaultLetClause) Let(aliases ...*Alias) WhereClause {
	p.setElement(&letElement{aliases})
	return newDefaultWhereClause(p)
}

func (p *defaultLetClause) Join(from interface{}) JoinClause {
	p.setElement(&joinElement{defaultJoin, toString(from)})
	return newDefaultJoinClause(p)
}

func (p *defaultLetClause) InnerJoin(from interface{}) JoinClause {
	p.setElement(&joinElement{inner, toString(from)})
	return newDefaultJoinClause(p)
}

func (p *defaultLetClause) LeftJoin(from interface{}) JoinClause {
	p.setElement(&joinElement{left, toString(from)})
	return newDefaultJoinClause(p)
}

func (p *defaultLetClause) LeftOuterJoin(from interface{}) JoinClause {
	p.setElement(&joinElement{leftOuter, toString(from)})
	return newDefaultJoinClause(p)
}

func (p *defaultLetClause) Nest(from interface{}) NestClause {
	p.setElement(&nestElement{defaultJoin, toString(from)})
	return newDefaultNestClause(p)
}

func (p *defaultLetClause) InnerNest(from interface{}) NestClause {
	p.setElement(&nestElement{inner, toString(from)})
	return newDefaultNestClause(p)
}

func (p *defaultLetClause) LeftNest(from interface{}) NestClause {
	p.setElement(&nestElement{left, toString(from)})
	return newDefaultNestClause(p)
}

func (p *defaultLetClause) LeftOuterNest(from interface{}) NestClause {
	p.setElement(&nestElement{leftOuter, toString(from)})
	return newDefaultNestClause(p)
}

func (p *defaultLetClause) Unnest(from interface{}) UnnestClause {
	p.setElement(newUnnestElement(defaultJoin, toString(from)))
	return newDefaultUnnestClause(p)
}

func (p *defaultLetClause) InnerUnnest(from interface{}) UnnestClause {
	p.setElement(newUnnestElement(inner, toString(from)))
	return newDefaultUnnestClause(p)
}

func (p *defaultLetClause) LeftUnnest(from interface{}) UnnestClause {
	p.setElement(newUnnestElement(left, toString(from)))
	return newDefaultUnnestClause(p)
}

func (p *defaultLetClause) LeftOuterUnnest(from interface{}) UnnestClause {
	p.setElement(newUnnestElement(leftOuter, toString(from)))
	return newDefaultUnnestClause(p)
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
