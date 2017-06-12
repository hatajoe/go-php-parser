package ast

import (
	"fmt"
)

type VariableExpression struct {
	Name Expression
	Prop Expression
}

func NewVariableExpression(name, prop Expression) *VariableExpression {
	return &VariableExpression{
		Name: name,
		Prop: prop,
	}
}

func (ve *VariableExpression) expressionNode()      {}
func (ve *VariableExpression) TokenLiteral() string { return "" }
func (ve *VariableExpression) String() string {
	return fmt.Sprintf("%s->%s", ve.Name.String(), ve.Prop.String())
}
