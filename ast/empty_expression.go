package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type EmptyExpression struct {
	Token *token.Token
	Arg   Expression
}

func NewEmptyExpression(tok *token.Token, arg Expression) *EmptyExpression {
	return &EmptyExpression{
		Token: tok,
		Arg:   arg,
	}
}

func (ee *EmptyExpression) expressionNode()      {}
func (ee *EmptyExpression) TokenLiteral() string { return ee.Token.Literal }
func (ee *EmptyExpression) String() string {
	return fmt.Sprintf("empty(%s)", ee.Arg.String())
}
