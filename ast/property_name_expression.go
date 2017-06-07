package ast

import (
	"fmt"
)

type PropertyNameExpression struct {
	Expr Expression
}

func NewPropertyNameExpression(expr Expression) *PropertyNameExpression {
	return &PropertyNameExpression{
		Expr: expr,
	}
}

func (de *PropertyNameExpression) expressionNode()      {}
func (de *PropertyNameExpression) TokenLiteral() string { return "" }
func (de *PropertyNameExpression) String() string       { return fmt.Sprintf("{%s}", de.Expr.String()) }
