package ast

import (
	"fmt"
)

type ArrayPairExpression struct {
	Left  Expression
	Right Expression
}

func NewArrayPairExpression(left, right Expression) *ArrayPairExpression {
	return &ArrayPairExpression{
		Left:  left,
		Right: right,
	}
}

func (ape *ArrayPairExpression) expressionNode()      {}
func (ape *ArrayPairExpression) TokenLiteral() string { return "" }
func (ape *ArrayPairExpression) String() string {
	return fmt.Sprintf("%s => %s", ape.Left.String(), ape.Right.String())
}
