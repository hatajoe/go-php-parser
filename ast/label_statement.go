package ast

import (
	"fmt"
)

type LabelStatement struct {
	Expr Expression
}

func NewLabelStatement(expr Expression) *LabelStatement {
	return &LabelStatement{
		Expr: expr,
	}
}

func (ls *LabelStatement) statementNode()       {}
func (ls *LabelStatement) TokenLiteral() string { return "" }
func (ls *LabelStatement) String() string {
	return fmt.Sprintf("%s:", ls.Expr.String())
}
