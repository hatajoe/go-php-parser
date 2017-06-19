package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type RequireExpression struct {
	Token *token.Token
	Expr  Expression
}

func NewRequireExpression(tok *token.Token, expr Expression) *RequireExpression {
	return &RequireExpression{
		Token: tok,
		Expr:  expr,
	}
}

func (re *RequireExpression) expressionNode()      {}
func (re *RequireExpression) TokenLiteral() string { return re.Token.Literal }
func (re *RequireExpression) String() string {
	return fmt.Sprintf("%s %s", re.TokenLiteral(), re.Expr.String())
}
