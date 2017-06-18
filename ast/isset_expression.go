package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type IssetExpression struct {
	Token *token.Token
	Args  []Expression
}

func NewIssetExpression(tok *token.Token, args ...Expression) *IssetExpression {
	return &IssetExpression{
		Token: tok,
		Args:  args,
	}
}

func (ie *IssetExpression) expressionNode()      {}
func (ie *IssetExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IssetExpression) String() string {
	exprs := make([]string, 0, len(ie.Args))
	for _, expr := range ie.Args {
		exprs = append(exprs, expr.String())
	}
	return fmt.Sprintf("isset(%s)", strings.Join(exprs, ", "))
}
