package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type ExitExpression struct {
	Token *token.Token
	Right Expression
}

func NewExitExpression(tok *token.Token, right Expression) *ExitExpression {
	return &ExitExpression{
		Token: tok,
		Right: right,
	}
}

func (ce *ExitExpression) expressionNode()      {}
func (ce *ExitExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *ExitExpression) String() string {
	if ce.Right == nil {
		return "exit"
	}
	return fmt.Sprintf("exit%s", ce.Right.String())
}
