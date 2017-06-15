package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type CloneExpression struct {
	Token *token.Token
	Right Expression
}

func NewCloneExpression(tok *token.Token, right Expression) *CloneExpression {
	return &CloneExpression{
		Token: tok,
		Right: right,
	}
}

func (ce *CloneExpression) expressionNode()      {}
func (ce *CloneExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CloneExpression) String() string {
	return fmt.Sprintf("clone %s", ce.Right.String())
}
