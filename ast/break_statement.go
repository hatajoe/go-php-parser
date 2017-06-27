package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type BreakStatement struct {
	Token *token.Token
	Expr  Expression
}

func NewBreakStatement(tok *token.Token, expr Expression) *BreakStatement {
	return &BreakStatement{
		Token: tok,
		Expr:  expr,
	}
}

func (bs *BreakStatement) statementNode()       {}
func (bs *BreakStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BreakStatement) String() string {
	if bs.Expr == nil {
		return fmt.Sprintf("%s;", bs.TokenLiteral())
	}
	return fmt.Sprintf("%s %s;", bs.TokenLiteral(), bs.Expr.String())
}
