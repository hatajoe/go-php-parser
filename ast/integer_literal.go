package ast

import (
	"strconv"

	"github.com/hatajoe/go-php-parser/token"
)

type IntegerLiteral struct {
	Token *token.Token
	Value int64
}

func NewIntegerLiteral(tok *token.Token, value string) *IntegerLiteral {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return &IntegerLiteral{
		Token: tok,
		Value: v,
	}
}

func (i *IntegerLiteral) expressionNode()      {}
func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *IntegerLiteral) String() string {
	return strconv.FormatInt(i.Value, 10)
}
