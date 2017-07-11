package ast

type OptionalTypeExpression struct {
	Expr Expression
}

func NewOptionalTypeExpression(expr Expression) *OptionalTypeExpression {
	return &OptionalTypeExpression{
		Expr: expr,
	}
}

func (ote *OptionalTypeExpression) expressionNode()      {}
func (ote *OptionalTypeExpression) TokenLiteral() string { return "" }
func (ote *OptionalTypeExpression) String() string {
	return "?" + ote.Expr.String()
}
