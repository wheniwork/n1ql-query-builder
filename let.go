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

	Join(from string) JoinPath

	InnerJoin(from string) JoinPath

	LeftJoin(from string) JoinPath

	LeftOuterJoin(from string) JoinPath

	Nest(from string) NestPath

	InnerNest(from string) NestPath

	LeftNest(from string) NestPath

	LeftOuterNest(from string) NestPath

	Unnest(from string) UnnestPath

	InnerUnnest(from string) UnnestPath

	LeftUnnest(from string) UnnestPath

	LeftOuterUnnest(from string) UnnestPath

	JoinExpr(from *Expression) JoinPath

	InnerJoinExpr(from *Expression) JoinPath

	LeftJoinExpr(from *Expression) JoinPath

	LeftOuterJoinExpr(from *Expression) JoinPath

	NestExpr(from *Expression) NestPath

	InnerNestExpr(from *Expression) NestPath

	LeftNestExpr(from *Expression) NestPath

	LeftOuterNestExpr(from *Expression) NestPath

	UnnestExpr(from *Expression) UnnestPath

	InnerUnnestExpr(from *Expression) UnnestPath

	LeftUnnestExpr(from *Expression) UnnestPath

	LeftOuterUnnestExpr(from *Expression) UnnestPath
}

type defaultLetPath struct {
	*defaultWherePath
}

func newDefaultLetPath(parent Path) *defaultLetPath {
	return &defaultLetPath{newDefaultWherePath(parent)}
}

func (p *defaultLetPath) Let(aliases ...*Alias) WherePath {
	p.setElement(newLetElement(aliases))
	return newDefaultWherePath(p)
}

func (p *defaultLetPath) Join(from string) JoinPath {
	p.setElement(newJoinElement(DefaultJoin, from))
	return newDefaultJoinPath(p)
}

func (p *defaultLetPath) InnerJoin(from string) JoinPath {
	p.setElement(newJoinElement(Inner, from))
	return newDefaultJoinPath(p)
}

func (p *defaultLetPath) LeftJoin(from string) JoinPath {
	p.setElement(newJoinElement(Left, from))
	return newDefaultJoinPath(p)
}

func (p *defaultLetPath) LeftOuterJoin(from string) JoinPath {
	p.setElement(newJoinElement(LeftOuter, from))
	return newDefaultJoinPath(p)
}

func (p *defaultLetPath) Nest(from string) NestPath {
	p.setElement(newNestElement(DefaultJoin, from))
	return newDefaultNestPath(p)
}

func (p *defaultLetPath) InnerNest(from string) NestPath {
	p.setElement(newNestElement(Inner, from))
	return newDefaultNestPath(p)
}

func (p *defaultLetPath) LeftNest(from string) NestPath {
	p.setElement(newNestElement(Left, from))
	return newDefaultNestPath(p)
}

func (p *defaultLetPath) LeftOuterNest(from string) NestPath {
	p.setElement(newNestElement(LeftOuter, from))
	return newDefaultNestPath(p)
}

func (p *defaultLetPath) Unnest(from string) UnnestPath {
	p.setElement(newUnnestElement(DefaultJoin, from))
	return newDefaultUnnestPath(p)
}

func (p *defaultLetPath) InnerUnnest(from string) UnnestPath {
	p.setElement(newUnnestElement(Inner, from))
	return newDefaultUnnestPath(p)
}

func (p *defaultLetPath) LeftUnnest(from string) UnnestPath {
	p.setElement(newUnnestElement(Left, from))
	return newDefaultUnnestPath(p)
}

func (p *defaultLetPath) LeftOuterUnnest(from string) UnnestPath {
	p.setElement(newUnnestElement(LeftOuter, from))
	return newDefaultUnnestPath(p)
}

func (p *defaultLetPath) JoinExpr(from *Expression) JoinPath {
	return p.Join(from.String())
}

func (p *defaultLetPath) InnerJoinExpr(from *Expression) JoinPath {
	return p.InnerJoin(from.String())
}

func (p *defaultLetPath) LeftJoinExpr(from *Expression) JoinPath {
	return p.LeftJoin(from.String())
}

func (p *defaultLetPath) LeftOuterJoinExpr(from *Expression) JoinPath {
	return p.LeftOuterJoin(from.String())
}

func (p *defaultLetPath) NestExpr(from *Expression) NestPath {
	return p.Nest(from.String())
}

func (p *defaultLetPath) InnerNestExpr(from *Expression) NestPath {
	return p.InnerNest(from.String())
}

func (p *defaultLetPath) LeftNestExpr(from *Expression) NestPath {
	return p.LeftNest(from.String())
}

func (p *defaultLetPath) LeftOuterNestExpr(from *Expression) NestPath {
	return p.LeftOuterNest(from.String())
}

func (p *defaultLetPath) UnnestExpr(from *Expression) UnnestPath {
	return p.Unnest(from.String())
}

func (p *defaultLetPath) InnerUnnestExpr(from *Expression) UnnestPath {
	return p.InnerUnnest(from.String())
}

func (p *defaultLetPath) LeftUnnestExpr(from *Expression) UnnestPath {
	return p.LeftUnnest(from.String())
}

func (p *defaultLetPath) LeftOuterUnnestExpr(from *Expression) UnnestPath {
	return p.LeftOuterUnnest(from.String())
}

type letElement struct {
	aliases []*Alias
}

func newLetElement(aliases []*Alias) *letElement {
	return &letElement{aliases}
}

func (e *letElement) Export() string {
	buf := bytes.NewBufferString("LET ")
	for i, alias := range e.aliases {
		buf.WriteString(alias.String())
		if i < len(e.aliases)-1 {
			buf.WriteString(", ")
		}
	}

	return buf.String()
}
