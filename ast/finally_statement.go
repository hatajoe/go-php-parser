package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type FinallyStatement struct {
	Token *token.Token
	Stmts []Statement
}

func NewFinallyStatement(tok *token.Token, stmts []Statement) *FinallyStatement {
	return &FinallyStatement{
		Token: tok,
		Stmts: stmts,
	}
}

func (fs *FinallyStatement) statementNode()       {}
func (fs *FinallyStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *FinallyStatement) String() string {
	stmts := make([]string, 0, len(fs.Stmts))
	for _, stmt := range fs.Stmts {
		stmts = append(stmts, "    "+stmt.String())
	}
	return fmt.Sprintf("%s {\n%s\n}", fs.TokenLiteral(), strings.Join(stmts, "\n"))
}
