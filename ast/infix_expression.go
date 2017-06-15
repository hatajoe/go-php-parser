package ast

import "fmt"

type InfixExpression struct {
	Type  InfixType
	Left  Expression
	Right Expression
}

func NewInfixExpression(t InfixType, left, right Expression) *InfixExpression {
	return &InfixExpression{
		Type:  t,
		Left:  left,
		Right: right,
	}
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return "" }
func (ie *InfixExpression) String() string {
	infix := ""
	switch ie.Type {
	case BooleanOr:
		infix = "||"
	case BooleanAnd:
		infix = "&&"
	case LogicalOr:
		infix = "or"
	case LogicalAnd:
		infix = "and"
	case LogicalXor:
		infix = "xor"
	case BwOr:
		infix = "|"
	case BwAnd:
		infix = "&"
	case BwXor:
		infix = "^"
	case Concat:
		infix = "."
	case Add:
		infix = "+"
	case Sub:
		infix = "-"
	case Mul:
		infix = "*"
	case Pow:
		infix = "**"
	case Div:
		infix = "/"
	case Mod:
		infix = "%"
	case Sl:
		infix = "<<"
	case Sr:
		infix = ">>"
	case IsIdentical:
		infix = "==="
	case IsNotIdentical:
		infix = "!=="
	case IsEqual:
		infix = "=="
	case IsNotEqual:
		infix = "!="
	case Smaller:
		infix = "<"
	case SmallerOrEqual:
		infix = "<="
	case Greater:
		infix = ">"
	case GreaterOrEqual:
		infix = ">="
	case Spaceship:
		infix = "<=>"
	case InstanceOf:
		infix = "instanceof"
	default:
		panic(fmt.Sprintf("err: Undefined type specified. got=%d", ie.Type))
	}
	return ie.Left.String() + " " + infix + " " + ie.Right.String()
}
