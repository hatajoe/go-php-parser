package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type YieldExpression struct {
	Token *token.Token
	Expr  Expression
}

func NewYieldExpression(tok *token.Token, expr Expression) *YieldExpression {
	return &YieldExpression{
		Token: tok,
		Expr:  expr,
	}
}

func (ye *YieldExpression) expressionNode()      {}
func (ye *YieldExpression) TokenLiteral() string { return ye.Token.Literal }
func (ye *YieldExpression) String() string {
	if ye.Expr == nil {
		return fmt.Sprintf("%s", ye.TokenLiteral())
	}
	return fmt.Sprintf("%s %s", ye.TokenLiteral(), ye.Expr.String())
}
