package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type AbstractLiteral struct {
	Token *token.Token
	Value string
}

func NewAbstractLiteral(tok *token.Token, value string) *AbstractLiteral {
	return &AbstractLiteral{
		Token: tok,
		Value: value,
	}
}

func (a *AbstractLiteral) expressionNode()      {}
func (a *AbstractLiteral) TokenLiteral() string { return a.Token.Literal }
func (a *AbstractLiteral) String() string       { return a.Value }
