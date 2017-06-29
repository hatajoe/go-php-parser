package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type TryStatement struct {
	Token   *token.Token
	Stmts   []Statement
	Catch   []Statement
	Finally Statement
}

func NewTryStatement(tok *token.Token, stmts []Statement, catch []Statement, finally Statement) *TryStatement {
	return &TryStatement{
		Token:   tok,
		Stmts:   stmts,
		Catch:   catch,
		Finally: finally,
	}
}

func (ts *TryStatement) statementNode()       {}
func (ts *TryStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *TryStatement) String() string {
	stmts := make([]string, 0, len(ts.Stmts))
	for _, stmt := range ts.Stmts {
		stmts = append(stmts, "    "+stmt.String())
	}
	catch := make([]string, 0, len(ts.Catch))
	for _, c := range ts.Catch {
		catch = append(catch, c.String())
	}
	if len(catch) > 0 {
		catch[0] = " " + catch[0]
	}
	finaly := ""
	if ts.Finally != nil {
		finaly = " " + ts.Finally.String()
	}
	return fmt.Sprintf("%s {\n%s\n}%s%s", ts.TokenLiteral(), strings.Join(stmts, "\n"), strings.Join(catch, " "), finaly)
}
