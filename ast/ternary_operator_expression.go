package ast

import "fmt"

type TernaryOperatorExpression struct {
	Condition   Expression
	Consequence Expression
	Alternative Expression
}

func NewTernaryOperatorExpression(cond, cons, alter Expression) *TernaryOperatorExpression {
	return &TernaryOperatorExpression{
		Condition:   cond,
		Consequence: cons,
		Alternative: alter,
	}
}

func (toe *TernaryOperatorExpression) expressionNode()      {}
func (toe *TernaryOperatorExpression) TokenLiteral() string { return "" }
func (toe *TernaryOperatorExpression) String() string {
	if toe.Consequence == nil {
		return fmt.Sprintf("%s ?: %s", toe.Condition, toe.Alternative)
	}
	return fmt.Sprintf("%s ? %s : %s", toe.Condition, toe.Consequence, toe.Alternative)
}
