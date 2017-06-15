package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type CastExpression struct {
	Token *token.Token
	Right Expression
}

func NewCastExpression(tok *token.Token, right Expression) *CastExpression {
	return &CastExpression{
		Token: tok,
		Right: right,
	}
}

func (ce *CastExpression) expressionNode()      {}
func (ce *CastExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CastExpression) String() string {
	return fmt.Sprintf("%s%s", ce.Token.Literal, ce.Right.String())
}
