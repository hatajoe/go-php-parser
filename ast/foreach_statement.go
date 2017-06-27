package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type ForeachStatement struct {
	Token *token.Token
	Expr  Expression
	Key   Expression
	Value Expression
	Stmt  Statement
}

func NewForeachStatement(tok *token.Token, expr, key, value Expression, stmt Statement) *ForeachStatement {
	return &ForeachStatement{
		Token: tok,
		Expr:  expr,
		Key:   key,
		Value: value,
		Stmt:  stmt,
	}
}

func (fs *ForeachStatement) statementNode()       {}
func (fs *ForeachStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForeachStatement) String() string {
	if fs.Key == nil {
		return fmt.Sprintf("%s (%s as %s) %s", fs.TokenLiteral(), fs.Expr.String(), fs.Value.String(), fs.Stmt.String())
	}
	return fmt.Sprintf("%s (%s as %s => %s) %s", fs.TokenLiteral(), fs.Expr.String(), fs.Key.String(), fs.Value.String(), fs.Stmt.String())
}
