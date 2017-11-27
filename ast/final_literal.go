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

func (f *FinalLiteral) expressionNode()      {}
func (f *FinalLiteral) TokenLiteral() string { return f.Token.Literal }
func (f *FinalLiteral) String() string       { return f.Value }
