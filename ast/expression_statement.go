package ast

import "fmt"

type ExpressionStatement struct {
	Expr Expression
}

func NewExpressionStatement(expr Expression) *ExpressionStatement {
	return &ExpressionStatement{
		Expr: expr,
	}
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return "" }
func (es *ExpressionStatement) String() string {
	return fmt.Sprintf("%s;", es.Expr.String())
}
