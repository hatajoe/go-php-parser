package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type PrivateLiteral struct {
	Token *token.Token
	Value string
}

func NewPrivateLiteral(tok *token.Token, value string) *PrivateLiteral {
	return &PrivateLiteral{
		Token: tok,
		Value: value,
	}
}

func (p *PrivateLiteral) expressionNode()      {}
func (p *PrivateLiteral) TokenLiteral() string { return p.Token.Literal }
func (p *PrivateLiteral) String() string       { return p.Value }
