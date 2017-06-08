package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type StaticLiteral struct {
	Token *token.Token
	Value string
}

func NewStaticLiteral(tok *token.Token, value string) *StaticLiteral {
	return &StaticLiteral{
		Token: tok,
		Value: value,
	}
}

func (s *StaticLiteral) expressionNode()      {}
func (s *StaticLiteral) TokenLiteral() string { return s.Token.Literal }
func (s *StaticLiteral) String() string       { return s.Value }
