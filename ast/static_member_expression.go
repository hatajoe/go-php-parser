package ast

import (
	"fmt"
)

type StaticMemberExpression struct {
	Name   Expression
	Member Expression
}

func NewStaticMemberExpression(name, member Expression) *StaticMemberExpression {
	return &StaticMemberExpression{
		Name:   name,
		Member: member,
	}
}

func (sme *StaticMemberExpression) expressionNode()      {}
func (sme *StaticMemberExpression) TokenLiteral() string { return "" }
func (sme *StaticMemberExpression) String() string {
	return fmt.Sprintf("%s::%s", sme.Name.String(), sme.Member.String())
}
