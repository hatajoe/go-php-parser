package ast

import (
	"fmt"
	"strings"
)

type ArrayExpression struct {
	Type  ArrayType
	Exprs []Expression
}

func NewArrayExpression(t ArrayType, exprs ...Expression) *ArrayExpression {
	return &ArrayExpression{
		Type:  t,
		Exprs: exprs,
	}
}

func (ae *ArrayExpression) expressionNode()      {}
func (ae *ArrayExpression) TokenLiteral() string { return "" }
func (ae *ArrayExpression) String() string {
	containsDoubleArrowPairs := false
	list := make([]string, 0, len(ae.Exprs))
	for _, expr := range ae.Exprs {
		if _, ok := expr.(*ArrayPairExpression); ok {
			containsDoubleArrowPairs = true
		}
		list = append(list, expr.String())
	}

	// Split array pair with 4 whitespaces and new line
	// when:
	//	- ae.Exprs contains double arrow pairs
	//	- ae.Exprs length more over 10
	//
	// e.g:
	//	array(
	//		"foo",
	//		"bar",
	//	);
	sep := ", "
	if containsDoubleArrowPairs || len(ae.Exprs) > 10 {
		list[0] = "\n    " + list[0]
		list[len(list)-1] = list[len(list)-1] + ",\n"
		sep = ",\n    "
	}

	switch ae.Type {
	case Long:
		return fmt.Sprintf("array(%s)", strings.Join(list, sep))
	case Short:
		return fmt.Sprintf("[%s]", strings.Join(list, sep))
	default:
		panic(fmt.Sprintf("err: Undefined type specified. got=%d", ae.Type))
	}
}
