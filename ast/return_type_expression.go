package ast

type ReturnTypeExpression struct {
	Expr Expression
}

func NewReturnTypeExpression(expr Expression) *ReturnTypeExpression {
	return &ReturnTypeExpression{
		Expr: expr,
	}
}

func (rte *ReturnTypeExpression) expressionNode()      {}
func (rte *ReturnTypeExpression) TokenLiteral() string { return "" }
func (rte *ReturnTypeExpression) String() string {
	return ": " + rte.Expr.String()
}
