package ast

import "fmt"

type DecrementExpression struct {
	Type  DecrementType
	Value Expression
}

func NewDecrementExpression(t DecrementType, value Expression) *DecrementExpression {
	return &DecrementExpression{
		Type:  t,
		Value: value,
	}
}

func (de *DecrementExpression) expressionNode()      {}
func (de *DecrementExpression) TokenLiteral() string { return "" }
func (de *DecrementExpression) String() string {
	switch de.Type {
	case PostDec:
		return fmt.Sprintf("%s--", de.Value.String())
	case PreDec:
		return fmt.Sprintf("--%s", de.Value.String())
	default:
		panic(fmt.Sprintf("err: Undefined type specified. got=%d", de.Type))
	}
}
