package ast

import (
	"fmt"
)

type CallableVariableExpression struct {
	Type  VariableType
	Value Expression
	Exprs []Expression
}

func NewCallableVariableExpression(t VariableType, val Expression, exprs ...Expression) *CallableVariableExpression {
	return &CallableVariableExpression{
		Type:  t,
		Value: val,
		Exprs: exprs,
	}
}

func (cve *CallableVariableExpression) expressionNode()      {}
func (cve *CallableVariableExpression) TokenLiteral() string { return "" }
func (cve *CallableVariableExpression) String() string {
	switch cve.Type {
	case Dim:
		expr := ""
		if cve.Exprs[0] != nil {
			expr = cve.Exprs[0].String()
		}
		return fmt.Sprintf("%s[%s]", cve.Value.String(), expr)
	case Curly:
		expr := ""
		if cve.Exprs[0] != nil {
			expr = cve.Exprs[0].String()
		}
		return fmt.Sprintf("%s{%s}", cve.Value.String(), expr)
	case Prop:
		return fmt.Sprintf("%s->%s%s", cve.Value.String(), cve.Exprs[0].String(), cve.Exprs[1].String())
	default:
		panic(fmt.Sprintf("err: Undefined type specified. got=%d", cve.Type))
	}
}
