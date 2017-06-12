package ast

import (
	"fmt"
)

type ConstantExpression struct {
	Name  Expression
	Ident Expression
}

func NewConstantExpression(name, ident Expression) *ConstantExpression {
	return &ConstantExpression{
		Name:  name,
		Ident: ident,
	}
}

func (ve *ConstantExpression) expressionNode()      {}
func (ve *ConstantExpression) TokenLiteral() string { return "" }
func (ve *ConstantExpression) String() string {
	return fmt.Sprintf("%s::%s", ve.Name.String(), ve.Ident.String())
}
