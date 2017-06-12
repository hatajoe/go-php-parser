package ast

import "fmt"

type AmpersandLiteral struct {
	Expr Expression
}

func NewAmpersandLiteral(expr Expression) *AmpersandLiteral {
	return &AmpersandLiteral{
		Expr: expr,
	}
}

func (al *AmpersandLiteral) expressionNode()      {}
func (al *AmpersandLiteral) TokenLiteral() string { return "" }
func (al *AmpersandLiteral) String() string {
	return fmt.Sprintf("&%s", al.Expr.String())
}
