package ast

import (
	"fmt"
)

type FunctionCallExpression struct {
	Type   FunctionCallType
	Name   Expression
	Member Expression
	Args   Expression
}

func NewFunctionCallExpression(t FunctionCallType, name, member, args Expression) *FunctionCallExpression {
	return &FunctionCallExpression{
		Type:   t,
		Name:   name,
		Member: member,
		Args:   args,
	}
}

func (fce *FunctionCallExpression) expressionNode()      {}
func (fce *FunctionCallExpression) TokenLiteral() string { return "" }
func (fce *FunctionCallExpression) String() string {
	switch fce.Type {
	case Call:
		return fmt.Sprintf("%s%s", fce.Name.String(), fce.Args.String())
	case StaticCall:
		return fmt.Sprintf("%s::%s%s", fce.Name.String(), fce.Member.String(), fce.Args.String())
	default:
		panic(fmt.Sprintf("err: Undefined type specified. got=%d", fce.Type))
	}
}
