package ast

import (
	"fmt"
)

type DereferencableExpression struct {
	Type  VariableType
	Exprs []Expression
}

func NewDereferencableExpression(t VariableType, exprs ...Expression) *DereferencableExpression {
	return &DereferencableExpression{
		Type:  t,
		Exprs: exprs,
	}
}

func (de *DereferencableExpression) expressionNode()      {}
func (de *DereferencableExpression) TokenLiteral() string { return "" }
func (de *DereferencableExpression) String() string {
	switch de.Type {
	case Dim:
		expr := ""
		if de.Exprs[1] != nil {
			expr = de.Exprs[1].String()
		}
		return fmt.Sprintf("%s[%s]", de.Exprs[0].String(), expr)
	case Wrapped:
		return fmt.Sprintf("(%s)", de.Exprs[0].String())
	default:
		panic(fmt.Sprintf("err: Undefined type specified. got=%d", de.Type))
	}
}
