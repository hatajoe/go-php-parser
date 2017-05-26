package ast

import (
	"strconv"

	"github.com/hatajoe/go-php-parser/token"
)

type DoubleLiteral struct {
	Token *token.Token
	Value float64
}

func NewDoubleLiteral(tok *token.Token, value string) *DoubleLiteral {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(err)
	}
	return &DoubleLiteral{
		Token: tok,
		Value: v,
	}
}

func (d *DoubleLiteral) expressionNode()      {}
func (d *DoubleLiteral) TokenLiteral() string { return d.Token.Literal }
func (d *DoubleLiteral) String() string       { return d.Token.Literal }
