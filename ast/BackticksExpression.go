package ast

import (
	"bytes"
)

type BackticksExpression struct {
	Exprs []Expression
}

func NewBackticksExpression(exprs ...Expression) *BackticksExpression {
	return &BackticksExpression{
		Exprs: exprs,
	}
}

func (ale *BackticksExpression) expressionNode()      {}
func (ale *BackticksExpression) TokenLiteral() string { return "" }
func (ale *BackticksExpression) String() string {
	var out bytes.Buffer

	out.WriteString("`")
	for _, expr := range ale.Exprs {
		out.WriteString(expr.String())
	}
	out.WriteString("`")

	return out.String()
}
