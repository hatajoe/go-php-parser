package ast

import (
	"fmt"

	"github.com/hatajoe/go-php-parser/token"
)

type SwitchStatement struct {
	Token *token.Token
	Cond  Expression
	Stmt  Statement
}

func NewSwitchStatement(tok *token.Token, cond Expression, stmt Statement) *SwitchStatement {
	return &SwitchStatement{
		Token: tok,
		Cond:  cond,
		Stmt:  stmt,
	}
}

func (ss *SwitchStatement) statementNode()       {}
func (ss *SwitchStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SwitchStatement) String() string {
	return fmt.Sprintf("%s (%s) %s", ss.TokenLiteral(), ss.Cond.String(), ss.Stmt.String())
}
