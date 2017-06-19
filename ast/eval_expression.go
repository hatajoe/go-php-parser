package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type EvalExpression struct {
	Token *token.Token
	Arg   Expression
}

func NewEvalExpression(tok *token.Token, arg Expression) *EvalExpression {
	return &EvalExpression{
		Token: tok,
		Arg:   arg,
	}
}

func (ee *EvalExpression) expressionNode()      {}
func (ee *EvalExpression) TokenLiteral() string { return ee.Token.Literal }
func (ee *EvalExpression) String() string {
	return fmt.Sprintf("eval(%s)", ee.Arg.String())
}
