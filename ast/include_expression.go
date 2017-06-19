package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type IncludeExpression struct {
	Token *token.Token
	Expr  Expression
}

func NewIncludeExpression(tok *token.Token, expr Expression) *IncludeExpression {
	return &IncludeExpression{
		Token: tok,
		Expr:  expr,
	}
}

func (ie *IncludeExpression) expressionNode()      {}
func (ie *IncludeExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IncludeExpression) String() string {
	return fmt.Sprintf("%s %s", ie.TokenLiteral(), ie.Expr.String())
}
