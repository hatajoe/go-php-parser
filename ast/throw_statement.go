package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type ThrowStatement struct {
	Token *token.Token
	Expr  Expression
}

func NewThrowStatement(tok *token.Token, expr Expression) *ThrowStatement {
	return &ThrowStatement{
		Token: tok,
		Expr:  expr,
	}
}

func (ts *ThrowStatement) statementNode()       {}
func (ts *ThrowStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *ThrowStatement) String() string {
	return fmt.Sprintf("%s %s;", ts.TokenLiteral(), ts.Expr.String())
}
