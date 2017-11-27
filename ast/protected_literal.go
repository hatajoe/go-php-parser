package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type ProtectedLiteral struct {
	Token *token.Token
	Value string
}

func NewProtectedLiteral(tok *token.Token, value string) *ProtectedLiteral {
	return &ProtectedLiteral{
		Token: tok,
		Value: value,
	}
}

func (p *ProtectedLiteral) expressionNode()      {}
func (p *ProtectedLiteral) TokenLiteral() string { return p.Token.Literal }
func (p *ProtectedLiteral) String() string       { return p.Value }
