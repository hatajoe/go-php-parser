package ast

import "fmt"

type IncrementExpression struct {
	Type  IncrementType
	Value Expression
}

func NewIncrementExpression(t IncrementType, value Expression) *IncrementExpression {
	return &IncrementExpression{
		Type:  t,
		Value: value,
	}
}

func (ie *IncrementExpression) expressionNode()      {}
func (ie *IncrementExpression) TokenLiteral() string { return "" }
func (ie *IncrementExpression) String() string {
	switch ie.Type {
	case PostInc:
		return fmt.Sprintf("%s++", ie.Value.String())
	case PreInc:
		return fmt.Sprintf("++%s", ie.Value.String())
	default:
		panic(fmt.Sprintf("err: Undefined type specified. got=%d", ie.Type))
	}
}
