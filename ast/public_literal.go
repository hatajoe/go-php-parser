package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type PublicLiteral struct {
	Token *token.Token
	Value string
}

func NewPublicLiteral(tok *token.Token, value string) *PublicLiteral {
	return &PublicLiteral{
		Token: tok,
		Value: value,
	}
}

func (p *PublicLiteral) expressionNode()      {}
func (p *PublicLiteral) TokenLiteral() string { return p.Token.Literal }
func (p *PublicLiteral) String() string       { return p.Value }
