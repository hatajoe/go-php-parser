package ast

import (
	"bytes"
	"strings"
)

type ArgumentListExpression struct {
	Exprs []Expression
}

func NewArgumentListExpression(exprs ...Expression) *ArgumentListExpression {
	return &ArgumentListExpression{
		Exprs: exprs,
	}
}

func (ale *ArgumentListExpression) expressionNode()      {}
func (ale *ArgumentListExpression) TokenLiteral() string { return "" }
func (ale *ArgumentListExpression) String() string {
	var out bytes.Buffer

	exprs := make([]string, 0, len(ale.Exprs))
	for _, expr := range ale.Exprs {
		exprs = append(exprs, expr.String())
	}

	out.WriteString("(")
	out.WriteString(strings.Join(exprs, ", "))
	out.WriteString(")")

	return out.String()
}
