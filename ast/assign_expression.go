package ast

import "fmt"

type AssignExpression struct {
	Type  AssignType
	Left  Expression
	Right Expression
	IsRef bool
}

func NewAssignExpression(t AssignType, left, right Expression, isRef bool) *AssignExpression {
	return &AssignExpression{
		Type:  t,
		Left:  left,
		Right: right,
		IsRef: isRef,
	}
}

func (ae *AssignExpression) expressionNode()      {}
func (ae *AssignExpression) TokenLiteral() string { return "" }
func (ae *AssignExpression) String() string {
	ref := ``
	if ae.IsRef {
		ref = `&`
	}

	assign := `=`
	switch ae.Type {
	case Equal:
		assign = `=`
	case PlusEqual:
		assign = `+=`
	case MinusEqual:
		assign = `-=`
	case MulEqual:
		assign = `*=`
	case PowEqual:
		assign = `**=`
	case DivEqual:
		assign = `/=`
	case ConcatEqual:
		assign = `.=`
	case ModEqual:
		assign = `%=`
	case AndEqual:
		assign = `&=`
	case QrEqual:
		assign = `|=`
	case XorEqual:
		assign = `^=`
	case SlEqual:
		assign = `<<=`
	case SrEqual:
		assign = `>>=`
	default:
		panic(fmt.Sprintf("err: Undefined type specified. got=%d", ae.Type))
	}

	return ae.Left.String() + " " + assign + " " + ref + ae.Right.String()
}
