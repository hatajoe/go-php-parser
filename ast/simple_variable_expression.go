package ast

import (
	"fmt"
)

type SimpleVariableExpression struct {
	Type VariableType
	Expr Expression
}

func NewSimpleVariableExpression(t VariableType, expr Expression) *SimpleVariableExpression {
	return &SimpleVariableExpression{
		Type: t,
		Expr: expr,
	}
}

func (ve *SimpleVariableExpression) expressionNode()      {}
func (ve *SimpleVariableExpression) TokenLiteral() string { return "" }
func (ve *SimpleVariableExpression) String() string {
	switch ve.Type {
	case Var:
		return fmt.Sprintf("$%s", ve.Expr.String())
	case CurlyOpen:
		return fmt.Sprintf("${%s}", ve.Expr.String())
	default:
		panic(fmt.Sprintf("err: Undefined type specified. got=%d", ve.Type))
	}
}
