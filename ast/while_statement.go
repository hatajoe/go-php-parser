package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type WhileStatement struct {
	Token *token.Token
	Expr  Expression
	Stmt  Statement
}

func NewWhileStatement(tok *token.Token, expr Expression, stmt Statement) *WhileStatement {
	return &WhileStatement{
		Token: tok,
		Expr:  expr,
		Stmt:  stmt,
	}
}

func (ws *WhileStatement) statementNode()       {}
func (ws *WhileStatement) TokenLiteral() string { return ws.Token.Literal }
func (ws *WhileStatement) String() string {
	return fmt.Sprintf("%s (%s) %s", ws.TokenLiteral(), ws.Expr.String(), ws.Stmt.String())
}
