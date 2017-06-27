package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type ContinueStatement struct {
	Token *token.Token
	Expr  Expression
}

func NewContinueStatement(tok *token.Token, expr Expression) *ContinueStatement {
	return &ContinueStatement{
		Token: tok,
		Expr:  expr,
	}
}

func (cs *ContinueStatement) statementNode()       {}
func (cs *ContinueStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ContinueStatement) String() string {
	if cs.Expr == nil {
		return fmt.Sprintf("%s;", cs.TokenLiteral())
	}
	return fmt.Sprintf("%s %s;", cs.TokenLiteral(), cs.Expr.String())
}
