package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type DoWhileStatement struct {
	Token *token.Token
	Expr  Expression
	Stmt  Statement
}

func NewDoWhileStatement(tok *token.Token, expr Expression, stmt Statement) *DoWhileStatement {
	return &DoWhileStatement{
		Token: tok,
		Expr:  expr,
		Stmt:  stmt,
	}
}

func (dws *DoWhileStatement) statementNode()       {}
func (dws *DoWhileStatement) TokenLiteral() string { return dws.Token.Literal }
func (dws *DoWhileStatement) String() string {
	return fmt.Sprintf("%s %s while (%s);", dws.TokenLiteral(), dws.Stmt.String(), dws.Expr.String())
}
