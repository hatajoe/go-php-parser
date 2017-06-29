package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type DeclareStatement struct {
	Token *token.Token
	Exprs []Expression
	Stmt  Statement
}

func NewDeclareStatement(tok *token.Token, exprs []Expression, stmt Statement) *DeclareStatement {
	return &DeclareStatement{
		Token: tok,
		Exprs: exprs,
		Stmt:  stmt,
	}
}

func (ds *DeclareStatement) statementNode()       {}
func (ds *DeclareStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DeclareStatement) String() string {
	exprs := make([]string, 0, len(ds.Exprs))
	for _, expr := range ds.Exprs {
		exprs = append(exprs, expr.String())
	}
	if _, ok := ds.Stmt.(*EmptyStatement); ok {
		return fmt.Sprintf("%s (%s)%s", ds.TokenLiteral(), strings.Join(exprs, ", "), ds.Stmt.String())
	}
	return fmt.Sprintf("%s (%s) %s", ds.TokenLiteral(), strings.Join(exprs, ", "), ds.Stmt.String())
}
