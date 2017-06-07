package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type ArgumentExpression struct {
	Token *token.Token
	Expr  Expression
}

func NewArgumentExpression(tok *token.Token, expr Expression) *ArgumentExpression {
	return &ArgumentExpression{
		Token: tok,
		Expr:  expr,
	}
}

func (ae *ArgumentExpression) expressionNode()      {}
func (ae *ArgumentExpression) TokenLiteral() string { return ae.Token.Literal }
func (ae *ArgumentExpression) String() string       { return fmt.Sprintf("...%s", ae.Expr.String()) }
