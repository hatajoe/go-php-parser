package ast

import (
	"fmt"
)

type NVariableExpression struct {
	Type  VariableType
	Value Expression
	Expr  Expression
}

func NewNVariableExpression(t VariableType, val, expr Expression) *NVariableExpression {
	return &NVariableExpression{
		Type:  t,
		Value: val,
		Expr:  expr,
	}
}

func (nve *NVariableExpression) expressionNode()      {}
func (nve *NVariableExpression) TokenLiteral() string { return "" }
func (nve *NVariableExpression) String() string {
	switch nve.Type {
	case Dim:
		return fmt.Sprintf("%s[%s]", nve.Value.String(), nve.Expr.String())
	case Curly:
		return fmt.Sprintf("%s{%s}", nve.Value.String(), nve.Expr.String())
	case Prop:
		return fmt.Sprintf("%s->%s", nve.Value.String(), nve.Expr.String())
	default:
		panic(fmt.Sprintf("err: Undefined type specified. got=%d", nve.Type))
	}
}
