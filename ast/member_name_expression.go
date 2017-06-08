package ast

import (
	"fmt"
)

type MemberNameExpression struct {
	Expr Expression
}

func NewMemberNameExpression(expr Expression) *MemberNameExpression {
	return &MemberNameExpression{
		Expr: expr,
	}
}

func (me *MemberNameExpression) expressionNode()      {}
func (me *MemberNameExpression) TokenLiteral() string { return "" }
func (me *MemberNameExpression) String() string       { return fmt.Sprintf("{%s}", me.Expr.String()) }
