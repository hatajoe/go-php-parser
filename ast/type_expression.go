package ast

import (
	"github.com/hatajoe/go-php-parser/token"
)

type TypeExpression struct {
	Token *token.Token
}

func NewTypeExpression(tok *token.Token) *TypeExpression {
	return &TypeExpression{
		Token: tok,
	}
}

func (te *TypeExpression) expressionNode()      {}
func (te *TypeExpression) TokenLiteral() string { return te.Token.Literal }
func (te *TypeExpression) String() string {
	return te.TokenLiteral()
}
