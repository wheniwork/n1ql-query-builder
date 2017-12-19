package nqb

import "encoding/json"

type KeysClauses interface {
	LetClause

	// OnKeys adds the on-key clause of a join/nest/unnest clause
	OnKeys(key interface{}) LetClause

	// OnKeysValues adds the on-key clause of a join/nest/unnest clause
	// with 1-n constant keys (eg. ON KEYS "a" or ON KEYS ["a", "b"])
	OnKeysValues(constantKeys ...string) LetClause

	// UseKeys sets the primary keyspace (doc id) in a join clause)
	UseKeys(key interface{}) LetClause

	// UseKeysValues sets the primary keyspace (doc id) in a join clause, with
	// one or more keys given as constants (eg. USE KEYS "test" or
	// USE KEYS ["a", "b"])
	UseKeysValues(keys ...string) LetClause
}

type defaultKeysClauses struct {
	*defaultLetClause
}

func newDefaultKeysClauses(parent Statement) *defaultKeysClauses {
	return &defaultKeysClauses{newDefaultLetClause(parent)}
}

func (p *defaultKeysClauses) OnKeys(key interface{}) LetClause {
	switch key.(type) {
	case *Expression:
		p.setElement(&keysElement{JoinOn, key.(*Expression)})
	default:
		p.setElement(&keysElement{JoinOn, X(key)})
	}

	return newDefaultLetClause(p)
}

func (p *defaultKeysClauses) OnKeysValues(constantKeys ...string) LetClause {
	if len(constantKeys) == 1 {
		return p.OnKeys(S(constantKeys[0]).String())
	}

	jsonBytes, err := json.Marshal(constantKeys)

	if err != nil {
		panic(err) //todo handle this better
	}

	return p.OnKeys(string(jsonBytes))
}

func (p *defaultKeysClauses) UseKeys(key interface{}) LetClause {
	switch key.(type) {
	case *Expression:
		p.setElement(&keysElement{UseKeyspace, key.(*Expression)})
	default:
		p.setElement(&keysElement{UseKeyspace, X(key)})
	}

	return newDefaultLetClause(p)
}

func (p *defaultKeysClauses) UseKeysValues(keys ...string) LetClause {
	if len(keys) == 1 {
		return p.UseKeys(S(keys[0]))
	}

	jsonBytes, err := json.Marshal(keys)

	if err != nil {
		panic(err) //todo handle this better
	}

	return p.UseKeys(string(jsonBytes))
}

type ClauseType string

const (
	JoinOn      ClauseType = "ON KEYS "
	UseKeyspace ClauseType = "USE KEYS "
)

type keysElement struct {
	clauseType ClauseType
	expression *Expression
}

func (e *keysElement) export() string {
	return string(e.clauseType) + e.expression.String()
}
