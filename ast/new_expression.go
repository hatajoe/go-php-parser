package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type NewExpression struct {
	Token *token.Token
	Class Expression
	Args  Expression
}

func NewNewExpression(tok *token.Token, class, args Expression) *NewExpression {
	return &NewExpression{
		Token: tok,
		Class: class,
		Args:  args,
	}
}

func (nne *NewExpression) expressionNode()      {}
func (nne *NewExpression) TokenLiteral() string { return nne.Token.Literal }
func (nne *NewExpression) String() string {
	if nne.Args == nil {
		return fmt.Sprintf("new %s", nne.Class.String())
	}
	return fmt.Sprintf("new %s%s", nne.Class.String(), nne.Args.String())
}
