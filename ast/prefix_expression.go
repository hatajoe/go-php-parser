package ast

import "fmt"

type PrefixExpression struct {
	Type  PrefixType
	Right Expression
}

func NewPrefixExpression(t PrefixType, right Expression) *PrefixExpression {
	return &PrefixExpression{
		Type:  t,
		Right: right,
	}
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return "" }
func (pe *PrefixExpression) String() string {
	prefix := ""
	switch pe.Type {
	case UnaryPlus:
		prefix = "+"
	case UnaryMinus:
		prefix = "-"
	case BoolNot:
		prefix = "!"
	case BwNot:
		prefix = "~"
	case Silence:
		prefix = "@"
	default:
		panic(fmt.Sprintf("err: Undefined type specified. got=%d", pe.Type))
	}
	return fmt.Sprintf("%s%s", prefix, pe.Right.String())
}
