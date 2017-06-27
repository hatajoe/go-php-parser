package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type ReturnStatement struct {
	Token *token.Token
	Expr  Expression
}

func NewReturnStatement(tok *token.Token, expr Expression) *ReturnStatement {
	return &ReturnStatement{
		Token: tok,
		Expr:  expr,
	}
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	if rs.Expr == nil {
		return fmt.Sprintf("%s;", rs.TokenLiteral())
	}
	return fmt.Sprintf("%s %s;", rs.TokenLiteral(), rs.Expr.String())
}
