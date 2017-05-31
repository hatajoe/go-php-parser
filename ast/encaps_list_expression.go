package ast

import (
	"bytes"
)

type EncapsListExpression struct {
	Exprs []Expression
}

func NewEncapsListExpression(exprs ...Expression) *EncapsListExpression {
	return &EncapsListExpression{
		Exprs: exprs,
	}
}

func (ele *EncapsListExpression) expressionNode()      {}
func (ele *EncapsListExpression) TokenLiteral() string { return `"` }
func (ele *EncapsListExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ele.TokenLiteral())
	for _, v := range ele.Exprs {
		out.WriteString(v.String())
	}
	out.WriteString(ele.TokenLiteral())

	return out.String()
}
