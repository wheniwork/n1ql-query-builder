package nqb

import "encoding/json"

type KeysPath interface {
	LetPath

	OnKeysExpr(expression *Expression) LetPath

	OnKeys(key string) LetPath

	OnKeysValues(constantKeys ...string) LetPath

	UseKeysExpr(expression *Expression) LetPath

	UseKeys(key string) LetPath

	UseKeysValues(keys ...string) LetPath
}

type defaultKeysPath struct {
	*defaultLetPath
}

func newDefaultKeysPath(parent Path) *defaultKeysPath {
	return &defaultKeysPath{newDefaultLetPath(parent)}
}

func (p *defaultKeysPath) OnKeysExpr(expression *Expression) LetPath {
	p.setElement(&keysElement{JoinOn, expression})
	return newDefaultLetPath(p)
}

func (p *defaultKeysPath) OnKeys(key string) LetPath {
	return p.OnKeysExpr(X(key))
}

func (p *defaultKeysPath) OnKeysValues(constantKeys ...string) LetPath {
	if len(constantKeys) == 1 {
		return p.OnKeys(S(constantKeys[0]).String())
	}

	jsonBytes, err := json.Marshal(constantKeys)

	if err != nil {
		panic(err) //todo handle this better
	}

	return p.OnKeys(string(jsonBytes))
}

func (p *defaultKeysPath) UseKeysExpr(expression *Expression) LetPath {
	p.setElement(&keysElement{UseKeyspace, expression})
	return newDefaultLetPath(p)
}

func (p *defaultKeysPath) UseKeys(key string) LetPath {
	return p.UseKeysExpr(X(key))
}

func (p *defaultKeysPath) UseKeysValues(keys ...string) LetPath {
	if len(keys) == 1 {
		return p.UseKeysExpr(S(keys[0]))
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
