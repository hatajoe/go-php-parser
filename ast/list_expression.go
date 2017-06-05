package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type ListExpression struct {
	Token *token.Token
	Exprs []Expression
}

func NewListExpression(tok *token.Token, exprs ...Expression) *ListExpression {
	return &ListExpression{
		Token: tok,
		Exprs: exprs,
	}
}

func (le *ListExpression) expressionNode()      {}
func (le *ListExpression) TokenLiteral() string { return le.Token.Literal }
func (le *ListExpression) String() string {
	list := make([]string, 0, len(le.Exprs))
	for _, expr := range le.Exprs {
		list = append(list, expr.String())
	}

	// Split list pair with 4 whitespaces and new line
	// when:
	//	- ae.Exprs length more over 10
	//
	// e.g:
	//	list(
	//		$foo,
	//		$bar,
	//	);
	sep := ", "
	if len(le.Exprs) > 10 {
		list[0] = "\n    " + list[0]
		list[len(list)-1] = list[len(list)-1] + ",\n"
		sep = ",\n    "
	}

	return fmt.Sprintf("list(%s)", strings.Join(list, sep))
}
