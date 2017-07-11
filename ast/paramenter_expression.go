package ast

import (
	"fmt"
)

type ParameterExpression struct {
	OptionalType Expression
	IsReference  int
	IsVariadic   int
	Variable     Expression
	Default      Expression
}

func NewParameterExpression(optional Expression, isReference, isVariadic int, variable, def Expression) *ParameterExpression {
	return &ParameterExpression{
		OptionalType: optional,
		IsReference:  isReference,
		IsVariadic:   isVariadic,
		Variable:     variable,
		Default:      def,
	}
}

func (pe *ParameterExpression) expressionNode()      {}
func (pe *ParameterExpression) TokenLiteral() string { return "" }
func (pe *ParameterExpression) String() string {
	optional := ""
	if pe.OptionalType != nil {
		optional = pe.OptionalType.String() + " "
	}
	ref := ""
	if pe.IsReference != 0 {
		ref = "&"
	}
	ellipsis := ""
	if pe.IsVariadic != 0 {
		ellipsis = "..."
	}
	if pe.Default != nil {
		return fmt.Sprintf("%s%s%s%s = %s", optional, ref, ellipsis, pe.Variable.String(), pe.Default.String())
	}
	return fmt.Sprintf("%s%s%s%s", optional, ref, ellipsis, pe.Variable.String())
}
