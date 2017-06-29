package ast

import (
	"fmt"
	"strings"

	"github.com/hatajoe/go-php-parser/token"
)

type CatchListStatement struct {
	Token *token.Token
	Class []Expression
	Value Expression
	Stmts []Statement
}

func NewCatchListStatement(tok *token.Token, class []Expression, value Expression, stmts []Statement) *CatchListStatement {
	return &CatchListStatement{
		Token: tok,
		Class: class,
		Value: value,
		Stmts: stmts,
	}
}

func (cls *CatchListStatement) statementNode()       {}
func (cls *CatchListStatement) TokenLiteral() string { return cls.Token.Literal }
func (cls *CatchListStatement) String() string {
	class := make([]string, 0, len(cls.Class))
	for _, c := range cls.Class {
		class = append(class, c.String())
	}
	stmts := make([]string, 0, len(cls.Stmts))
	for _, stmt := range cls.Stmts {
		stmts = append(stmts, "    "+stmt.String())
	}
	return fmt.Sprintf("%s (%s %s) {\n%s\n}", cls.TokenLiteral(), strings.Join(class, "|"), cls.Value.String(), strings.Join(stmts, "\n"))
}
