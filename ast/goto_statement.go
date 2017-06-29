package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type GotoStatement struct {
	Token *token.Token
	Expr  Expression
}

func NewGotoStatement(tok *token.Token, expr Expression) *GotoStatement {
	return &GotoStatement{
		Token: tok,
		Expr:  expr,
	}
}

func (gs *GotoStatement) statementNode()       {}
func (gs *GotoStatement) TokenLiteral() string { return gs.Token.Literal }
func (gs *GotoStatement) String() string {
	return fmt.Sprintf("%s %s;", gs.TokenLiteral(), gs.Expr.String())
}
