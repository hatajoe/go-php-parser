package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type PrintExpression struct {
	Token *token.Token
	Right Expression
}

func NewPrintExpression(tok *token.Token, right Expression) *PrintExpression {
	return &PrintExpression{
		Token: tok,
		Right: right,
	}
}

func (pe *PrintExpression) expressionNode()      {}
func (pe *PrintExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrintExpression) String() string {
	return fmt.Sprintf("print %s", pe.Right.String())
}
