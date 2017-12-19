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

func (c *defaultLetClause) Let(aliases ...*Alias) WhereClause {
	c.setElement(&letElement{aliases})
	return newDefaultWhereClause(c)
}

func (c *defaultLetClause) Join(from interface{}) JoinClause {
	c.setElement(&joinElement{defaultJoin, toString(from)})
	return newDefaultJoinClause(c)
}

func (c *defaultLetClause) InnerJoin(from interface{}) JoinClause {
	c.setElement(&joinElement{inner, toString(from)})
	return newDefaultJoinClause(c)
}

func (c *defaultLetClause) LeftJoin(from interface{}) JoinClause {
	c.setElement(&joinElement{left, toString(from)})
	return newDefaultJoinClause(c)
}

func (c *defaultLetClause) LeftOuterJoin(from interface{}) JoinClause {
	c.setElement(&joinElement{leftOuter, toString(from)})
	return newDefaultJoinClause(c)
}

func (c *defaultLetClause) Nest(from interface{}) NestClause {
	c.setElement(&nestElement{defaultJoin, toString(from)})
	return newDefaultNestClause(c)
}

func (c *defaultLetClause) InnerNest(from interface{}) NestClause {
	c.setElement(&nestElement{inner, toString(from)})
	return newDefaultNestClause(c)
}

func (c *defaultLetClause) LeftNest(from interface{}) NestClause {
	c.setElement(&nestElement{left, toString(from)})
	return newDefaultNestClause(c)
}

func (c *defaultLetClause) LeftOuterNest(from interface{}) NestClause {
	c.setElement(&nestElement{leftOuter, toString(from)})
	return newDefaultNestClause(c)
}

func (c *defaultLetClause) Unnest(from interface{}) UnnestClause {
	c.setElement(newUnnestElement(defaultJoin, toString(from)))
	return newDefaultUnnestClause(c)
}

func (c *defaultLetClause) InnerUnnest(from interface{}) UnnestClause {
	c.setElement(newUnnestElement(inner, toString(from)))
	return newDefaultUnnestClause(c)
}

func (c *defaultLetClause) LeftUnnest(from interface{}) UnnestClause {
	c.setElement(newUnnestElement(left, toString(from)))
	return newDefaultUnnestClause(c)
}

func (c *defaultLetClause) LeftOuterUnnest(from interface{}) UnnestClause {
	c.setElement(newUnnestElement(leftOuter, toString(from)))
	return newDefaultUnnestClause(c)
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
