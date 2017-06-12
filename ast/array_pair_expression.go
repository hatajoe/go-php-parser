package ast

import (
	"fmt"
)

type ArrayPairExpression struct {
	Left  Expression
	Right Expression
	IsRef bool
}

func NewArrayPairExpression(left, right Expression, isRef bool) *ArrayPairExpression {
	return &ArrayPairExpression{
		Left:  left,
		Right: right,
		IsRef: isRef,
	}
}

func (ape *ArrayPairExpression) expressionNode()      {}
func (ape *ArrayPairExpression) TokenLiteral() string { return "" }
func (ape *ArrayPairExpression) String() string {
	if ape.IsRef {
		return fmt.Sprintf("%s => &%s", ape.Left.String(), ape.Right.String())
	}
	return fmt.Sprintf("%s => %s", ape.Left.String(), ape.Right.String())
}
