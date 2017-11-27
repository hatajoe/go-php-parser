package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type FinalLiteral struct {
	Token *token.Token
	Value string
}

func NewFinalLiteral(tok *token.Token, value string) *FinalLiteral {
	return &FinalLiteral{
		Token: tok,
		Value: value,
	}
}

func (a *FinalLiteral) expressionNode()      {}
func (a *FinalLiteral) TokenLiteral() string { return a.Token.Literal }
func (a *FinalLiteral) String() string       { return a.Value }
