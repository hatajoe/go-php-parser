package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type AltWhileStatement struct {
	Token *token.Token
	Stmts []Statement
}

func NewAltWhileStatement(tok *token.Token, stmts ...Statement) *AltWhileStatement {
	return &AltWhileStatement{
		Token: tok,
		Stmts: stmts,
	}
}

func (aws *AltWhileStatement) statementNode()       {}
func (aws *AltWhileStatement) TokenLiteral() string { return aws.Token.Literal }
func (aws *AltWhileStatement) String() string {
	stmts := make([]string, 0, len(aws.Stmts))
	for _, stmt := range aws.Stmts {
		stmts = append(stmts, "    "+stmt.String())
	}
	return fmt.Sprintf(":\n%s\n%s;", strings.Join(stmts, "\n"), aws.TokenLiteral())
}
